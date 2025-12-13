package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []byte(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

func TestPart1(t *testing.T) {
	lines := bytes.Split(input, []byte("\n"))
	output := part1(lines)
	assert.Equal(t, 13, output)
}

func TestPart2(t *testing.T) {
	lines := bytes.Split(input, []byte("\n"))
	output := part2(lines)
	assert.Equal(t, 43, output)
}
