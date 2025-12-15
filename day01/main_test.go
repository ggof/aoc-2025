package main

import (
	"testing"

	"ggof.xyz/aoc2025/lib"
	"github.com/stretchr/testify/assert"
)

func TestDay01Part1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	lines := lib.Lines(input)
	res := part1(lines)

	assert.Equal(t, 3, res)
}
