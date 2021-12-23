package math2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGF28_AddSub(t *testing.T) {
	a := []GF28{GF28(0x98), GF28(0x10), GF28(0x90), GF28(0x5C), GF28(0x3F)}
	b := []GF28{GF28(0x88), GF28(0x8), GF28(0xC4), GF28(0xA7)}
	for i := 1; i < 5; i++ {
		assert.Equal(t, b[i-1], a[0].AddSub(a[i]))
	}
}

func TestGF28_Times(t *testing.T) {
	a := []GF28{GF28(0x9C), GF28(0x23), GF28(0x46), GF28(0x8C), GF28(0x3), 0x6, 0xC, 0x18}
	for i := 1; i < 8; i++ {
		assert.Equal(t, a[i], Times(a[i-1]))
	}
}

func TestGF28_Order(t *testing.T) {
	a := []GF28{GF28(0x9C), GF28(0x23), GF28(0x46), GF28(0x8C), GF28(0x3), GF28(0x0), GF28(0x1)}
	b := []int{7, 5, 6, 7, 1, 0, 0}
	for i, o := range b {
		assert.Equal(t, o, a[i].Order(), "p(x)=%x", a[i])
	}
}

func TestGF28_Multiply(t *testing.T) {
	a := []GF28{0x12, 0x43, 0xcc, 0x5d}
	b := []GF28{0xaa, 0x67, 0x1b, 0xec}
	for i := range a {
		assert.Equal(t, a[i].Multiply(b[i]), GF28(0x1), "a[i]:%x b[i]:%x", a[i], b[i])
	}
	assert.Equal(t, GF28(0x31).Multiply(GF28(0x81)), GF28(0x1))
}

func TestPolyDivide(t *testing.T) {
	assert.Equal(t, GF28(0xc), GF28(0x14).PolyDivide(0x3))
	assert.Equal(t, GF28(0x4), GF28(0x31).PolyDivide(0xc))
	assert.Equal(t, GF28(0x0), GF28(0x3).PolyDivide(0x18))

}

func TestGF28_Inverse(t *testing.T) {
	a := []GF28{0xFF, 0x11, 0x1a}
	b := []GF28{0x1c, 0xb4, 0xfd}
	for i := range a {
		assert.Equal(t, a[i].Inverse(), b[i])
	}
}
