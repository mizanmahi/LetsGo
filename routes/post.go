package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PostRoutes(r chi.Router) {
	r.Get("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome from post routes"))
	})
}