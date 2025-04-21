// internal/user/usecase/interface.go
package usecase

import "chi-project/internal/user"

type UserUsecase interface {
	CreateUser(u *user.User) (int, error)
	GetAllUsers() ([]*user.User, error)
}
