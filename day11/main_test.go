package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestPart1(t *testing.T) {
	result := part1([]byte(testInput))
	assert.Equal(t, 5, result)
}

func TestAppend(t *testing.T) {
	a := []string{"a"}
	t.Log(a)
	b := append(a, "b")
	t.Log(a, b)
	c := append(a, "c")
	t.Log(a, b, c)

	t.Fail()
}
