package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	output := part1(strings.Split(input, "\n"))
	assert.Equal(t, 357, output)
}

func TestPart2(t *testing.T) {
	output := part2(strings.Split(input, "\n"))
	assert.Equal(t, 3121910778619, output)
}
