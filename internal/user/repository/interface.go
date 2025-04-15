package repository

import "os/user"

type UserRepository interface {
	CreateUser(user *user.User) (int, error)
	GetUserByID(id int) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]*user.User, error)
}