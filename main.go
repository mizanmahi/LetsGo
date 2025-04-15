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
	_ "github.com/lib/pq"
)

const dbString = "postgres://postgres:541990@localhost:5432/gocrud?sslmode=disable"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func main() {

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}


	config.LoadEnvs()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routes.UserRoutes(router)
	routes.PostRoutes(router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Create table if not exists
	// createTable(db)

	// Create a new user
	// createUser(db, User{
	// 	Username: "john_doe",
	// 	Email:    "john@example.com",
	// 	Password: "password123",
	// })

	// getUser(db)
	users, err := getAllUsers(db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Users:", users)

	// Start Server
	port := config.AppConfig.Port
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	

}

// Create table if not exists, generally its done in migrations
// func createTable(db *sql.DB) {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		username VARCHAR(255) NOT NULL,
// 		email VARCHAR(255) NOT NULL,
// 		password VARCHAR(255) NOT NULL
// 	);`

// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func createUser (db *sql.DB, user User) (int, error) {

// 	query := `insert into users (username, email, password) values ($1, $2, $3) returning id`

// 	var id int

// 	err := db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&id)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func getUser(db *sql.DB) (User, error) {
// 	var username, email, password string

// 	query := `select username, email, password  from users where  id = $1`
// 	err := db.QueryRow(query, 1).Scan(&username, &email, &password)

// 	if err != nil {
// 		return User{}, err
// 	}

// 	fmt.Println(username, email, password)

// 	return User{
// 		Username: username,
// 		Email:    email,
// 		Password: password,
// 	}, nil

// }

func getAllUsers (db *sql.DB) ([]User, error) {
	var users []User
	query := `select username, email, password  from users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}