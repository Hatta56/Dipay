package main

import (
	"Dipay/configs"
	"Dipay/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.CompanyRoute(router)
	routes.AdminRoute(router)
	routes.EmployeeRoute(router)

	fmt.Println("Listening on port 8000")
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatal(err)
	}

}
