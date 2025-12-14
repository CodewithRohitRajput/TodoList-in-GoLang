package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func TodoRoutes(router *mux.Router){
	router.HandleFunc("/create" , controllers.Create).Methods("POST")
	router.HandleFunc("/todos" , controllers.GetAll).Methods("GET")
}	