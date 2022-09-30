package usecase

import (
	"errors"

	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/repository"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
)

type usecase struct {
	repository repository.Repository
}

func (u *usecase) Login(email *valueobject.Email, password string) (*entity.User, error) {
	user, err := u.repository.Get(email)
	if err != nil {
		return nil, err
	}

	if user.Password.Verify(password) {
		return user, errors.New("failed to verify the password")
	}
	return user, nil
}

func (u *usecase) Register(
	email *valueobject.Email,
	name *valueobject.Name,
	password *valueobject.Password,
) error {
	user := entity.NewUser(email, name, password)
	if err := u.repository.Create(user); err != nil {
		return err
	}
	return nil
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{repository: repository}
}
