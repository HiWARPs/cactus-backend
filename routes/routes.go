package routes

import (
	"net/http"

	"github.com/HiWARPs/cactus-backend/controllers"
	"github.com/HiWARPs/cactus-backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", controllers.Health())

	r.Get("/project", controllers.GetAllProjects())
	r.Get("/project/{pid}", controllers.GetProject())
	r.Get("/project/{pid}/form/{id}", controllers.GetForm())

	r.Post("/signup", controllers.Signup())
	r.Post("/login", controllers.Login())
	r.Get("/logout", controllers.Logout())

	r.Post("/file", controllers.UploadFile())
	r.Post("/query_electrons", controllers.QueryElectrons())
	r.Get("/aggregate/{id}", controllers.AgrregateElectrons())

	r.Mount("/", adminRouter())

	return r
}

func adminRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(AdminOnly)

	r.Post("/file", controllers.UploadFile())
	r.Get("/electron", controllers.GetElectron())
	r.Get("/electronID", controllers.GetElectronID())
	r.Post("/electron", controllers.AddElectron())
	r.Put("/electronID", controllers.EditElectronID())

	r.Route("/project", func(r chi.Router) {
		r.Post("/", controllers.AddProject())
		r.Delete("/{pid}", controllers.DeleteProject())
		r.Put("/{pid}", controllers.UpdateProject())
	})

	r.Route("/project/{pid}/form", func(r chi.Router) {
		r.Post("/", controllers.AddForm())
		r.Delete("/{id}", controllers.DeleteForm())
		r.Put("/{id}", controllers.UpdateForm())
	})

	return r
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if utils.IsAuthorized(w, r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		return
	})
}
