package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestPart1(t *testing.T) {
	output := part1(strings.Split(input, "\n"))
	assert.Equal(t, 50, output)
}
