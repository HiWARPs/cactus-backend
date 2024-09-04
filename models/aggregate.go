package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AggregateQuery struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type AggregateRequest struct {
	ID      	string 			 `json:"id"`

}

type AggregateData struct {
	ID   primitive.ObjectID   `json:"_id" bson:"_id"`
	Data []map[string][]float64 `json:"data" bson:"data"`
}

