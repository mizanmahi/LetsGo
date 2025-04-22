package repository

import "chi-project/internal/user"

type UserRepository interface {
	CreateUser(user *user.User) (int, error)
	GetAllUsers() ([]*user.User, error)
	UpdateUser(u *user.User) error
	// DeleteUser(id int) error
}