package repository

import "chi-project/internal/user"

type UserRepository interface {
	CreateUser(user *user.User) (int, error)
	GetAllUsers() ([]*user.User, error)
	// GetUserByID(id int) (*user.User, error)
	// UpdateUser(user *user.User) error
	// DeleteUser(id int) error
}