package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	n1 := NewName("John", "Doe")
	n2 := NewName("John", "")

	assert.Equal(t, "John Doe", n1.Full())
	assert.Equal(t, "John", n2.Full())
}
