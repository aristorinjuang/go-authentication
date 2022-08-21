package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	password := NewPassword("")

	password.SetHash("password")

	assert.True(t, password.Verify("password"))
}
