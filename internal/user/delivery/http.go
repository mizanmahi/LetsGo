package delivery

import (
	"chi-project/internal/user"
	"chi-project/internal/user/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler (userUseCase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUseCase ,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user_id ,err := h.UserUsecase.CreateUser(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "User created with ID: %d", user_id)
		return
	}


	w.WriteHeader(http.StatusCreated)
}


func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUsecase.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}