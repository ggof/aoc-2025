package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"ggof.xyz/aoc2025/lib"
)

type point struct {
	x, y int
}

var surroundings = []point{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}

func removableRolls(lines [][]byte) []point {
	ps := make([]point, 0)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '@' {
				continue
			}

			var as int
			for _, p := range surroundings {
				if y+p.y < 0 || y+p.y >= len(lines) {
					continue
				}

				if x+p.x < 0 || x+p.x >= len(lines[y]) {
					continue
				}

				if lines[y+p.y][x+p.x] == '@' {
					as++
				}

			}

			if as < 4 {
				ps = append(ps, point{x, y})
			}
		}
	}

	return ps
}

func removeRolls(lines [][]byte, ps []point) {
	for _, p := range ps {
		lines[p.y][p.x] = '.'
	}
}

func part1(lines [][]byte) int {
	rollsPos := removableRolls(lines)
	return len(rollsPos)
}

func part2(lines [][]byte) int {
	var total int

	for {
		rolls := removableRolls(lines)
		if len(rolls) == 0 {
			break
		}

		total += len(rolls)

		removeRolls(lines, rolls)
	}

	return total
}

type Runner[T any] interface {
	run() T
}

func main() {
	file := lib.Must(os.Open("inputs/day04.txt"))
	input := lib.Must(io.ReadAll(file))
	lines := bytes.Split(input, []byte("\n"))

	start := time.Now()
	p1 := part1(lines)
	done1 := time.Now()
	p2 := part2(lines)
	done2 := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, done1.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, done2.Sub(done1))
}
