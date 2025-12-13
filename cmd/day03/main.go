package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"ggof.xyz/aoc2025/lib"
)

func loop(l string, size int) int {
	lenl := len(l)

	runes := make([]rune, size)

	first := 0
	for i := range runes {
		last := lenl + i + 1 - size

		var jj int
		for j, p := range l[first:last] {
			if p > runes[i] {
				runes[i] = p
				jj = j
			}
		}

		first += jj + 1
	}

	return lib.Must(strconv.Atoi(string(runes)))
}

func part1(input []string) int {
	var total int
	for _, l := range input {
		total += loop(l, 2)
	}

	return total
}

func part2(input []string) int {
	total := 0
	for _, l := range input {
		total += loop(l, 12)
	}

	return total
}

func main() {
	file := lib.Must(os.Open("inputs/day03.txt"))
	input := lib.Must(io.ReadAll(file))
	lines := lib.Lines(strings.Trim(string(input), "\n"))
	fmt.Printf("day01/part1: %d\n", part1(lines))
	fmt.Printf("day01/part2: %d\n", part2(lines))
}
