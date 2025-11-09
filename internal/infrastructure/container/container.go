package container

import (
	"github.com/denerkrebs/WorkflowEngine/config"
	"github.com/denerkrebs/WorkflowEngine/internal/application/usecase"
	domainRepo "github.com/denerkrebs/WorkflowEngine/internal/domain/repository"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/database"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB

	UserRepo            domainRepo.UserRepository
	RegisterUserUseCase *usecase.RegisterUserUseCase
}

func New(cfg *config.Config) (*Container, error) {
	db, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(db); err != nil {
		panic(err)
	}

	// Repositorios
	userRepo := repository.NewUserRepository(db)

	// Use cases ou services
	createUserUseCase := usecase.NewRegisterUser(userRepo)

	return &Container{
		DB:                  db,
		UserRepo:            userRepo,
		RegisterUserUseCase: createUserUseCase,
	}, nil
}
