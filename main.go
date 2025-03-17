package main

import (
	"chi-project/config"
	"chi-project/database"
	"chi-project/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {

	config.LoadEnvs()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routes.UserRoutes(router)
	routes.PostRoutes(router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Connect database
	database.ConnectDB()


	// Start Server
	port := config.AppConfig.Port
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	

}
