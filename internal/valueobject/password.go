package valueobject

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Hash string
}

func (p *Password) SetHash(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	p.Hash = string(hash)
}

func (p *Password) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(password)) == nil
}

func NewPassword(hash string) *Password {
	return &Password{
		Hash: hash,
	}
}
