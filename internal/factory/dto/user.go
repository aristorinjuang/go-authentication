package dto

import (
	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
)

type User struct {
	Email     *valueobject.Email
	FirstName string
	LastName  string
	Hash      string
}

func (u *User) Entity() *entity.User {
	return entity.NewUser(
		u.Email,
		valueobject.NewName(u.FirstName, u.LastName),
		valueobject.NewPasswordFromHash(u.Hash),
	)
}

func NewUser(email *valueobject.Email) *User {
	return &User{
		Email: email,
	}
}
