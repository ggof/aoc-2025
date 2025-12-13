package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMod(t *testing.T) {
	a := -100
	b := 100

	c := Mod(a, b)

	assert.Equal(t, 0, c)
}
