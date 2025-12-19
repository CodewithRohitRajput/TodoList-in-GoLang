package main

import (
	"log"
	"net/http"
	"server/config"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
	"server/routes"

	"github.com/gorilla/mux"
)

func main(){
	config.ConnectDB()

	router := mux.NewRouter();
	routes.TodoRoutes(router)

	srv := &http.Server{
		Addr : ":8000",
		Handler : router,
	}

	go func(){
		log.Println("Server is running on port 8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit , syscall.SIGINT , syscall.SIGTERM)
	<- quit
	log.Println("Shutting down server...")

	cts,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(cts); err != nil{
		log.Fatal("Server failed to shutdown:", err)
	}
	
	log.Println("Server exited properly")

}