package model

import (
	"time"

	"github.com/denerkrebs/WorkflowEngine/internal/domain/entity"
)

type User struct {
	ID          string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Email       string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Login       string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password    string     `gorm:"not null"`
	Active      bool       `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"default:null"`
	LastLoginAt *time.Time `gorm:"default:null"`
}

func FromEntity(e *entity.User) *User {
	return &User{
		ID:          e.ID,
		Name:        e.Name,
		Email:       e.Email,
		Login:       e.Password,
		Password:    e.Password,
		Active:      e.Active,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		LastLoginAt: e.LastLoginAt,
	}
}

func (m *User) ToEntity() *entity.User {
	return &entity.User{
		ID:          m.ID,
		Name:        m.Name,
		Email:       m.Email,
		Login:       m.Password,
		Password:    m.Password,
		Active:      m.Active,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		LastLoginAt: m.LastLoginAt,
	}
}
