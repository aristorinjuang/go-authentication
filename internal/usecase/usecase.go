package usecase

import (
	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
)

type Usecase interface {
	Login(email *valueobject.Email, password string) (*entity.User, error)
	Register(*valueobject.Email, *valueobject.Name, *valueobject.Password) error
}
