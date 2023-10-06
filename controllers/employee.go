package controllers

import (
	"Dipay/configs"
	"Dipay/models"
	"Dipay/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var employeeCollection *mongo.Collection = configs.GetCollection(configs.DB, "employee")

func GetEmployeeByCompanyID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		employeeID := mux.Vars(r)["id"]
		defer cancel()

		var employee models.Employee
		err := employeeCollection.FindOne(ctx, bson.M{"company_id": employeeID}).Decode(&employee)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employee}}
		json.NewEncoder(rw).Encode(response)
	}

}

func GetEmployeeID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		employeeID := mux.Vars(r)["id"]
		defer cancel()

		var employee models.Employee
		err := employeeCollection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&employee)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employee}}
		json.NewEncoder(rw).Encode(response)
	}

}

func CreateEmployee() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var employee models.Employee
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&employee); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		result, err := employeeCollection.InsertOne(ctx, employee)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.EmployeeResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)

	}
}

func UpdateEmployee() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var employee models.Employee
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&employee); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		result, err := employeeCollection.UpdateOne(ctx, bson.M{"_id": employee.ID}, bson.M{"$set": employee})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func DeleteEmployee() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		employeeID := mux.Vars(r)["id"]

		result, err := employeeCollection.DeleteOne(ctx, bson.M{"_id": employeeID})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}
