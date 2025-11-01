package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID          string
	Name        string
	Email       string
	Login       string
	Password    string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	LastLoginAt *time.Time
}

type NewUserParams struct {
	Name     string
	Email    string
	Login    string
	Password string
}

func NewUser(params NewUserParams) (*User, error) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Name, validation.Required),
		validation.Field(&params.Email, validation.Required, is.Email),
		validation.Field(&params.Login, validation.Required),
		validation.Field(&params.Password, validation.Required),
	)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:      params.Name,
		Email:     params.Email,
		Login:     params.Login,
		Password:  params.Password,
		Active:    true,
		CreatedAt: time.Now(),
	}, nil
}
