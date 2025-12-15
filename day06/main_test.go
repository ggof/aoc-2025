package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `)

func TestPart1(t *testing.T) {
	output := Part1(input)
	assert.Equal(t, 4277556, output)
}

func TestPart2(t *testing.T) {
	output := Part2(input)
	assert.Equal(t, 3263827, output)
}
