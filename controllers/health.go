package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderOutput := render.New()
		renderOutput.JSON(w, 200, nil)
	}
}
