package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Completed bool               `json:"completed"`
}
