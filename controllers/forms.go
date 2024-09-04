package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HiWARPs/cactus-backend/models"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		var form models.Form
		json.NewDecoder(r.Body).Decode(&form)
		form.ID = primitive.NewObjectID()

		pid := chi.URLParam(r, "pid")
		projectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		_, err := mh.UpdateProject(bson.D{{"_id", projectId}}, bson.M{"$push": bson.M{"forms": form}})

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 201, map[string]string{"message": "form added successfully"})
	}
}

func GetForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		projectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		id := chi.URLParam(r, "id")
		formId, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		project := &models.Project{}
		err := mh.GetOneProject(project, bson.M{"_id": projectId})
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		foundForm, err := findForm(formId, project.Forms)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, foundForm)
	}
}

func findForm(id primitive.ObjectID, list []models.Form) (models.Form, error) {
	var foundForm models.Form
	for _, form := range list {
		if form.ID == id {
			foundForm = form
			break
		}
	}
	if foundForm.ID.IsZero() {
		return foundForm, fmt.Errorf("Missing id")
	}
	return foundForm, nil
}

func DeleteForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		projectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		id := chi.URLParam(r, "id")
		formId, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		_, err := mh.UpdateProject(bson.D{{"_id", projectId}},
			bson.M{"$pull": bson.M{"forms": bson.M{"_id": formId}}})

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]string{"message": "form deleted successfully"})
	}
}

func UpdateForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "pid")
		projectId, idErr := primitive.ObjectIDFromHex(pid)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}
		filter := bson.D{{"_id", projectId}}

		id := chi.URLParam(r, "id")
		formId, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		var changes map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
			http.Error(w, fmt.Sprintf("failed to parses changes needed %s", err), 500)
			return
		}

		project := &models.Project{}
		err := mh.GetOneProject(project, filter)
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		newForms := changeForms(formId, project.Forms, changes)
		_, err = mh.UpdateProject(filter, bson.D{{"$set", bson.M{"forms": newForms}}})
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]string{"message": "form updated successfully"})
	}
}

func changeForms(id primitive.ObjectID, list []models.Form, changes map[string]interface{}) []models.Form {
	for i, form := range list {
		if form.ID == id {
			if changes["name"] != nil {
				list[i].Name = changes["name"].(string)
			}
			if changes["description"] != nil {
				list[i].Description = changes["description"].(string)
			}
			if changes["references"] != nil {
				list[i].References = changes["references"].(string)
			}
			if changes["range_variable"] != nil {
				list[i].RangeVariable = changes["range_variable"].(string)
			}
			if changes["hidden_variables"] != nil {
				list[i].HiddenVariables = convertToSliceString(changes["hidden_variables"])
			}
			break
		}
	}
	return list
}

func convertToSliceString(value interface{}) []string {
	v := value.([]interface{})
	result := make([]string, len(v))
	for i, el := range v {
		str, ok := el.(string)
		if !ok {
			return []string{}
		}
		result[i] = str
	}
	return result
}
