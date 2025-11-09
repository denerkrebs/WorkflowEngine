package repository

import (
	"context"

	"github.com/denerkrebs/WorkflowEngine/internal/domain/entity"
	"github.com/denerkrebs/WorkflowEngine/internal/domain/repository"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(ctx context.Context, user *entity.User) error {
	model := model.FromEntity(user)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	*user = *model.ToEntity()
	return nil
}
