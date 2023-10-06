package routes

import (
	"Dipay/controllers"

	"github.com/gorilla/mux"
)

func CompanyRoute(router *mux.Router) {
	router.HandleFunc("/api/companies", controllers.CreateCompanies()).Methods("POST")
	router.HandleFunc("/api/companies", controllers.GetAllCompany()).Methods("GET")
	router.HandleFunc("/api/companies/{companyId}/set_active", controllers.SetActiveCompany()).Methods("PUT")
}
