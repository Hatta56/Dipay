package controllers

import (
	"Dipay/configs"
	"Dipay/models"
	"Dipay/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var companiesCollection *mongo.Collection = configs.GetCollection(configs.DB, "companies")
var validate = validator.New()

func CreateCompanies() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var company models.Company
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&company); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newCompany := models.Company{
			ID:              primitive.NewObjectID(),
			CompanyName:     company.CompanyName,
			TelephoneNumber: company.TelephoneNumber,
			IsActive:        company.IsActive,
			Address:         company.Address,
		}

		result, err := companiesCollection.InsertOne(ctx, newCompany)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.CompanyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func SetActiveCompany() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		companyID := params["CompanyID"]
		var company models.Company
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(companyID)

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&company); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{"is_active": company.IsActive}

		result, err := companiesCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//get updated company
		var updateCompany models.Company
		if result.MatchedCount == 1 {
			err := companiesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateCompany)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateCompany}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllCompany() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var company []models.Company
		defer cancel()

		results, err := companiesCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCompany models.Company
			if err = results.Decode(&singleCompany); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}

			company = append(company, singleCompany)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": company}}
		json.NewEncoder(rw).Encode(response)
	}
}
