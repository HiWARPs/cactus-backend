package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	References      string             `json:"references" bson:"references"`
	RangeVariable   string             `json:"range_variable" bson:"range_variable"`
	HiddenVariables []string           `json:"hidden_variables" bson:"hidden_variables"`
}
