package repository

import (
	"context"

	"github.com/denerkrebs/WorkflowEngine/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}
