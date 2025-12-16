package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"server/config"
	"server/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.Completed = false

	_, err := config.TodoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.TodoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo
	cursor.All(context.Background(), &todos)
	json.NewEncoder(w).Encode(todos)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	objectId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.TodoCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Successfully deleted")

}


func GetById (w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var todo models.Todo
	objectId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	err = config.TodoCollection.FindOne(context.Background() , bson.M{"_id" :objectId}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo);

}


func Update (w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")

	vars := mux.Vars(r)
	id  := vars["id"]

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	update := bson.M{
		"$set": bson.M{
			"title" : todo.Title,
			"completed" : todo.Completed,
		},
	}

	objectId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.TodoCollection.UpdateOne(context.Background() , bson.M{"_id" : objectId} , update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Successfull updated bruh")
}