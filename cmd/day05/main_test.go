package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

func TestPart1(t *testing.T) {
	output := part1(input)
	assert.Equal(t, 3, output)
}
