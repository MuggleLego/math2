package math2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMillerRabin(t *testing.T) {
	assert.Equal(t, MillerRabin(1000000009, 10), true)
	assert.Equal(t, MillerRabin(2, 10), true)
	assert.Equal(t, MillerRabin(5001, 10), false)
	assert.Equal(t, MillerRabin(561, 10), false)
}
