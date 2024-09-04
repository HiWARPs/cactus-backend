package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/unrolled/render"
)

type ElectronRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ElectronIDRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetElectron() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]bool{"success": true})
	}
}

func AddElectron() http.HandlerFunc {
	renderOutput := render.New()

	return func(w http.ResponseWriter, r *http.Request) {
		var req ElectronRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			renderOutput.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		result := 0

		renderOutput.JSON(w, 200, map[string]int{"result": result})
	}
}

func GetElectronID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]bool{"success": true})
	}
}

func EditElectronID() http.HandlerFunc {
	renderOutput := render.New()

	return func(w http.ResponseWriter, r *http.Request) {
		var req ElectronRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			renderOutput.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		result := 0

		renderOutput.JSON(w, 200, map[string]int{"result": result})
	}
}
