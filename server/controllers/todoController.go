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
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.Completed = false

	_, err := config.TodoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Println("Error inserting todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create todo"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := config.TodoCollection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Println("Error finding todos:" , err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error":"Failed to fetch todos"})
		return
	}
	defer cursor.Close(context.Background());


	
	var todos []models.Todo
 if err :=	cursor.All(context.Background(), &todos); err != nil {
	log.Println("Error decoding todos:" , err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error":"Failed to decode todos"})
	return
 }
	json.NewEncoder(w).Encode(todos)
}




func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	objectId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error" : "Invalid Id"})
		return
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid Id"})
		return
	}

	err = config.TodoCollection.FindOne(context.Background() , bson.M{"_id" :objectId}).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error" : "Todo not found"})
		return
	}

	
	json.NewEncoder(w).Encode(todo);
}


func Update (w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type" , "application/json")

	vars := mux.Vars(r)
	id  := vars["id"]

	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid todo data"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title" : todo.Title,
			"completed" : todo.Completed,
		},
	}

	objectId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid Id"})
			return
	}


	result, err := config.TodoCollection.UpdateOne(context.Background() , bson.M{"_id" : objectId} , update)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":"Failed to update todo"})
			return
	}
	json.NewEncoder(w).Encode(result)
	}
