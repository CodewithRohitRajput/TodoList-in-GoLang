package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func TodoRoutes(router *mux.Router) {
	router.HandleFunc("/create", controllers.Create).Methods("POST")
	router.HandleFunc("/todos", controllers.GetAll).Methods("GET")
	router.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	router.HandleFunc("/update/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.GetById).Methods("GET")

}
