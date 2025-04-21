package main

import (
	"chi-project/internal/config"
	"chi-project/internal/db"
	"chi-project/internal/routes"
	"chi-project/internal/user/delivery"
	"chi-project/internal/user/repository"
	"chi-project/internal/user/usecase"

	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)




func main() {


	config.LoadEnvs()
	db := db.InitPostgres()
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userRepo := repository.NewUserRepo(db)
	userUseCase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewUserHandler(userUseCase)


	routes.RegisterUserRoutes(r, userHandler)


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Start Server
	port := config.AppConfig.Port
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}



