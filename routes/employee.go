package routes

import (
	"Dipay/controllers"

	"github.com/gorilla/mux"
)

func EmployeeRoute(router *mux.Router) {
	router.HandleFunc("/api/employees/{employeeId}", controllers.GetEmployeeID()).Methods("GET")
	router.HandleFunc("/api/employees/{employeeId}", controllers.UpdateEmployee()).Methods("PUT")
	router.HandleFunc("/api/employees/{employeeId}", controllers.DeleteEmployee()).Methods("DELETE")
	router.HandleFunc("/api/companies/{companyId}/employees", controllers.GetEmployeeByCompanyID()).Methods("GET")
	router.HandleFunc("/api/companies/{companyId}/employees", controllers.CreateEmployee()).Methods("POST")
}
