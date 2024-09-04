package controllers

import (
	"fmt"
	"net/http"

	"github.com/HiWARPs/cactus-backend/models"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AgrregateElectrons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		id := chi.URLParam(r, "id")
		query := models.AggregateRequest{}

		objectId, idErr := primitive.ObjectIDFromHex(id)

		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		electron, err := mh.AggregateElectron(query, bson.M{"_id": objectId})
		if err != nil {
			print("aaaaa")
			http.Error(w, fmt.Sprint(err), 500)
			return
		}
		renderOutput := render.New()
		renderOutput.JSON(w, 200, electron)
	}
}
