package controllers

import (
	"Dipay/configs"
	"Dipay/models"
	"Dipay/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admin")

func Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admin models.Admin
		defer cancel()

		if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		filter := bson.D{{"username", admin.Username}}
		var foundAdmin models.Admin
		err := adminCollection.FindOne(ctx, filter).Decode(&foundAdmin)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if foundAdmin.Password != admin.Password {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.AdminResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "incorrect password"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": foundAdmin}}
		json.NewEncoder(rw).Encode(response)
	}

}
