// internal/routes/routes.go
package routes

import (
	"chi-project/internal/user/delivery"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, h *delivery.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.GetAllUsers)
		r.Put("/{id}", h.UpdateUser)
	})
}
