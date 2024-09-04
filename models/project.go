package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Forms  []Form             `json:"forms" bson:"forms"`
	FileID primitive.ObjectID `json:"file_id" bson:"file_id"`
}

type ProjectCreate struct {
	Name string `json:"name" bson:"name"`
}
