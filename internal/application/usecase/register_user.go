package usecase

import (
	"context"

	"github.com/denerkrebs/WorkflowEngine/internal/application/dto"
	"github.com/denerkrebs/WorkflowEngine/internal/domain/entity"
	"github.com/denerkrebs/WorkflowEngine/internal/domain/repository"
)

type RegisterUserUseCase struct {
	userRepo repository.UserRepository
}

func NewRegisterUser(userRepo repository.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, userDto dto.UserDto) error {
	user, err := entity.NewUser(entity.NewUserParams{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Login:    userDto.Login,
		Password: userDto.Password,
	})

	if err != nil {
		return err
	}

	uc.userRepo.Create(ctx, user)

	return nil
}
