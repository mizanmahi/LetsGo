package main

import (
	"chi-project/config"
	"chi-project/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const dbString = "postgres://postgres:postgres@localhost:5432/gocrud?sslmode=disable"


func main() {

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	config.LoadEnvs()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routes.UserRoutes(router)
	routes.PostRoutes(router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})


	// Start Server
	port := config.AppConfig.Port
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	

}
