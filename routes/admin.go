package routes

import (
	"Dipay/controllers"

	"github.com/gorilla/mux"
)

func AdminRoute(router *mux.Router) {
	router.HandleFunc("/api/admins/login", controllers.Login()).Methods("POST")
}
