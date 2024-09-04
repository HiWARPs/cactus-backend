package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RangeQuery struct {
	Name       string  `json:"name"`
	LowerBound float64 `json:"lower_bound"`
	UpperBound float64 ` json:"upper_bound"`
	Increment  float64 `json:"increment"`
}

type XQuery struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Request struct {
	ID      string `json:"id"`
	Range   RangeQuery
	X       []XQuery `json:"x"`
	Exclude []string `json:"exclude"`
}

func (r *Request) IsValid(query Request) error {
	for _, xq := range query.X {
		if xq.Name == query.Range.Name {
			return fmt.Errorf("multiple X queries detected")
		}
	}
	return nil
}

type ElectronData struct {
	ID   primitive.ObjectID   `json:"_id" bson:"_id"`
	Data []map[string]float64 `json:"data" bson:"data"`
}
