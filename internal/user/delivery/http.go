package delivery

import (
	"chi-project/internal/user"
	"chi-project/internal/user/usecase"
	"chi-project/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	user_id ,err := h.UserUsecase.CreateUser(&u); 
	 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "User created with ID: %d", user_id)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]int{"user_id": user_id})
	
}


func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUsecase.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateUser called")
	var u user.User

	user_id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err :=  strconv.Atoi(user_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ID = id

	err = h.UserUsecase.UpdateUser(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "User updated successfully"})	


}