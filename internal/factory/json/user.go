package json

import (
	"time"

	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
	"github.com/golang-jwt/jwt/v4"
)

type accessToken struct {
	AccessToken string `json:"access_token"`
}

type token struct {
	accessToken
	RefreshToken string `json:"refresh_token"`
}

type tokenSecret struct {
	access  []byte
	refresh []byte
}

type name struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Full  string `json:"full"`
}

type User struct {
	tokenSecret tokenSecret
	Email       string `json:"email"`
	Name        name   `json:"name"`
	jwt.RegisteredClaims
}

func (u *User) AccessToken() *accessToken {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	atSS, _ := at.SignedString(u.tokenSecret.access)

	return &accessToken{atSS}
}

func (u *User) Entity() *entity.User {
	return entity.NewUser(
		valueobject.NewEmail(u.Email),
		valueobject.NewName(u.Name.First, u.Name.Last),
		new(valueobject.Password),
	)
}

func (u *User) Token() *token {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	atSS, _ := at.SignedString(u.tokenSecret.access)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	rtSS, _ := rt.SignedString(u.tokenSecret.refresh)

	return &token{
		accessToken:  accessToken{atSS},
		RefreshToken: rtSS,
	}
}

func NewUser(user *entity.User, accessTokenDuration time.Duration, accessTokenSecret, refreshTokenSecret string) *User {
	return &User{
		tokenSecret: tokenSecret{
			access:  []byte(accessTokenSecret),
			refresh: []byte(refreshTokenSecret),
		},
		Email: user.Email.String(),
		Name: name{
			First: user.Name.First,
			Last:  user.Name.Last,
			Full:  user.Name.Full(),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		},
	}
}
