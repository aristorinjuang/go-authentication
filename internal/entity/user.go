package entity

import "github.com/aristorinjuang/go-authentication/internal/valueobject"

type User struct {
	Email    *valueobject.Email
	Name     *valueobject.Name
	Password *valueobject.Password
}

func NewUser(email *valueobject.Email, name *valueobject.Name, password *valueobject.Password) *User {
	return &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
}
