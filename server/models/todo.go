package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type Todo struct{
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Completed bool `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}