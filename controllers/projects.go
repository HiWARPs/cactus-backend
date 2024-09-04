package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/HiWARPs/cactus-backend/models"
	"github.com/HiWARPs/cactus-backend/mongo"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mh *mongo.MongoHandler

func Connect() error {
	// Замість завантаження конфігурації з файлу, отримуємо значення з оточення
	mongoDbConnection := os.Getenv("MONGO_URL")
	if mongoDbConnection == "" {
		return errors.New("MONGO_URL environment variable is not set")
	}

	var err error
	mh, err = mongo.NewHandler(mongoDbConnection)

	return err
}

func AddProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		project := &models.ProjectCreate{}
		json.NewDecoder(r.Body).Decode(&project)

		_, err := mh.AddOneProject(project)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 201, map[string]string{"message": "project added successfully"})
	}
}

func GetAllProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		projects, err := mh.GetProjects(bson.M{})
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}
		renderOutput := render.New()
		renderOutput.JSON(w, 200, projects)
	}
}

func GetProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		project := &models.Project{}
		objectId, idErr := primitive.ObjectIDFromHex(pid)

		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		err := mh.GetOneProject(project, bson.M{"_id": objectId})

		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, project)
	}
}

func DeleteProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		objectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		project := &models.Project{}
		err := mh.GetOneProject(project, bson.M{"_id": objectId})
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}
		_, err = mh.RemoveOneProject(bson.M{"_id": objectId})
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]string{"message": "project deleted successfully"})
	}
}

func UpdateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		objectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}
		filter := bson.D{{"_id", objectId}}

		var changes map[string]interface{}
		json.NewDecoder(r.Body).Decode(&changes)

		if fileIDStr, ok := changes["file_id"].(string); ok {
			fileID, idErr := primitive.ObjectIDFromHex(fileIDStr)
			if idErr != nil {
				http.Error(w, fmt.Sprintf("Invalid file_id"), 400)
				return
			}
			changes["file_id"] = fileID
		}

		update := bson.D{{"$set", changes}}

		_, err := mh.UpdateProject(filter, update)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]string{"message": "project updated successfully"})
	}
}
