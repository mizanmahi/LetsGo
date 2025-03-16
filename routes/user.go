package routes

import (
	"chi-project/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router *chi.Mux) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/", controllers.GetUsers) 
	})
}