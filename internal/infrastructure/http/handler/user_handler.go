package handler

import (
	"encoding/json"
	"net/http"

	"github.com/denerkrebs/WorkflowEngine/internal/application/dto"
	"github.com/denerkrebs/WorkflowEngine/internal/domain/entity"
)

// definir a dependencia
type UserHandler struct {
}

// injetar a dependencia
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	json.NewDecoder(r.Body).Decode(&userDto)

	user, err := entity.NewUser(entity.NewUserParams{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Login:    userDto.Login,
		Password: userDto.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
