package main

import (
	"log"
	"net/http"
	"server/config"

	"server/routes"

	"github.com/gorilla/mux"
)

func main(){
	config.ConnectDB()

	router := mux.NewRouter();
	routes.TodoRoutes(router)

	log.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", router)

}