package valueobject

import "strings"

type Email struct {
	Local  string
	Domain string
}

func (e *Email) String() string {
	return e.Local + "@" + e.Domain
}

func NewEmail(email string) *Email {
	emailChunks := strings.Split(email, "@")
	local := strings.Join(emailChunks[:len(emailChunks)-1], "@")
	domain := emailChunks[len(emailChunks)-1]

	return &Email{
		Local:  local,
		Domain: domain,
	}
}
