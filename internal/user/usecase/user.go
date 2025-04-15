package usecase

import (
	"chi-project/internal/user"
	"chi-project/internal/user/repository"
)

type userUseCase struct {
	repo repository.UserRepository
}



func NewUserUsecase(repo repository.UserRepository) *userUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(u *user.User) (int, error) {
	// Business logic here if needed
	return uc.repo.CreateUser(u)
}

func (uc *userUseCase) GetAllUsers() ([]*user.User, error) {
	return uc.repo.GetAllUsers()
}


// router => delivery => usecase (its interface will be used in delivery) => repository (its interface will be used in usecase)