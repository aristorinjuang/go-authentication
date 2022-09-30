package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	t.Run("set hash", func(t *testing.T) {
		password := new(Password)

		password.SetHash("password")

		assert.True(t, password.Verify("password"))
	})

	t.Run("from plain", func(t *testing.T) {
		password := NewPasswordFromPlain("password")

		assert.True(t, password.Verify("password"))
	})

	t.Run("from hash", func(t *testing.T) {
		password := NewPasswordFromHash("$2a$10$Q597AWauDGPoEOYiRzIF6.3Oti.r3GOfD0tUsKRpg24R7GOMdIXY.")

		assert.Equal(t, "$2a$10$Q597AWauDGPoEOYiRzIF6.3Oti.r3GOfD0tUsKRpg24R7GOMdIXY.", password.Hash())
	})
}
