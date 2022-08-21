package repository

import (
	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
)

type Repository interface {
	Get(*valueobject.Email) (*entity.User, error)
	Create(*entity.User) error
}
