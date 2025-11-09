package handler

import (
	"encoding/json"
	"net/http"

	"github.com/denerkrebs/WorkflowEngine/internal/application/dto"
	"github.com/denerkrebs/WorkflowEngine/internal/application/usecase"
)

type UserHandler struct {
	registerUserUseCase *usecase.RegisterUserUseCase
}

func NewUserHandler(registerUserUseCase usecase.RegisterUserUseCase) *UserHandler {
	return &UserHandler{
		registerUserUseCase: &registerUserUseCase,
	}
}

func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	json.NewDecoder(r.Body).Decode(&userDto)

	err := h.registerUserUseCase.Execute(r.Context(), userDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
