package http

import (
	"encoding/json"
	"movie-festival/internal/entity"
	"movie-festival/internal/usecase"
	"net/http"
)

type UserHandler struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: uc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    user := &entity.User{}
    if err := json.NewDecoder(r.Body).Decode(user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    userID, err := h.Usecase.Register(user)
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    id := int(userID)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]int{"userID": id})
}


func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	loggedInUser, err := h.Usecase.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loggedInUser)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := h.Usecase.Logout()
	if err != nil {
		http.Error(w, "Failed to logout user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged out successfully"})
}