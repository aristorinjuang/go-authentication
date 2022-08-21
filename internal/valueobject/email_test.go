package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	email := NewEmail("test@example.com")

	assert.Equal(t, "test@example.com", email.String())
}
