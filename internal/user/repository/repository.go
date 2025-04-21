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
func (r *userRepo) CreateUser(user *user.User) (int, error) {
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

func (r *userRepo) GetAllUsers() ([]*user.User, error) {
	rows, err := r.DB.Query(`SELECT id, username, email, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// explain below code
	// 1. Create a slice of pointers to user.User structs to hold the results.
	// 2. Iterate over the rows returned by the query using rows.Next().
	// 3. For each row, create a new user.User struct and scan the values from the row into the struct fields.
	// 4. Append the pointer to the user.User struct to the slice of users.
	// 5. Return the slice of user.User pointers and any error that occurred during the process.
	var users []*user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}