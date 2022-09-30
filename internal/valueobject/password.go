package valueobject

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func (p Password) Hash() string {
	return p.hash
}

func (p *Password) SetHash(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	p.hash = string(hash)
}

func (p *Password) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(password)) == nil
}

func NewPasswordFromHash(hash string) *Password {
	return &Password{
		hash: hash,
	}
}

func NewPasswordFromPlain(plainPassword string) *Password {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	return NewPasswordFromHash(string(hash))
}
