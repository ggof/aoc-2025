package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"ggof.xyz/aoc2025/lib"
)

var ops = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

func transpose(input [][]string) [][]string {
	output := make([][]string, len(input[0]))
	for _, as := range input {
		for j, a := range as {
			output[j] = append(output[j], a)
		}
	}
	return output
}

func Part1(input []byte) int {
	linesBytes := bytes.Split(input, []byte("\n"))

	lines := make([][]string, len(linesBytes))
	for i, l := range linesBytes {
		for f := range bytes.FieldsSeq(l) {
			lines[i] = append(lines[i], string(f))
		}
	}

	lines = transpose(lines)

	var total int
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		last := len(l) - 1
		op := l[last]

		lt := lib.Must(strconv.Atoi(l[0]))
		for _, n := range l[1:last] {
			nb := lib.Must(strconv.Atoi(n))
			lt = ops[op](lt, nb)
		}
		total += lt
	}

	return total
}

func Part2(input []byte) int {
	lines := strings.Split(string(input), "\n")
	nblines := len(lines)
	linelen := len(lines[0])
	lines, opsLine := lines[:nblines-1], lines[nblines-1]
	opsSlice := strings.Fields(opsLine)

	var iop, total int
	var nbs []int
	for i := range linelen {
		var nbBytes []byte
		for j := range len(lines) {
			nbBytes = append(nbBytes, lines[j][i])
		}

		nb := strings.TrimSpace(string(nbBytes))
		if len(nb) == 0 {
			op, acc := ops[opsSlice[iop]], nbs[0]
			for _, nb := range nbs[1:] {
				acc = op(acc, nb)
			}
			total += acc
			nbs = nil
			iop++
		} else {
			nbs = append(nbs, lib.Must(strconv.Atoi(nb)))
		}
	}

	op := ops[opsSlice[iop]]
	acc := nbs[0]
	for _, nb := range nbs[1:] {
		acc = op(acc, nb)
	}
	return total + acc
}

func main() {
	file := lib.Must(os.Open("inputs/day06.txt"))
	bs := lib.Must(io.ReadAll(file))
	start := time.Now()
	p1 := Part1(bs)
	done1 := time.Now()
	p2 := Part2(bs)
	done2 := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, done1.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, done2.Sub(done1))
}
