package repository

import (
	"chi-project/internal/user"
	"database/sql"
)

type userRepo struct {
	DB *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepo (db *sql.DB) *userRepo {
	return &userRepo{
		DB: db,
	}
}


// CreateUser creates a new user in the database
func (r *userRepo) CreateUser(user user.User) (int, error) {
	var userId int
	query := `insert into users (username, email, password) values ($1, $2, $3) returning id`

	err := r.DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *userRepo) GetAllUsers() ([]user.User, error) {
	rows, err := r.DB.Query(`SELECT id, username, email, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}