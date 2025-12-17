package main

import (
	"bytes"
	"fmt"
	"os"

	"ggof.xyz/aoc2025/lib"
)

type pos struct {
	x, y int
}

func part1(lines [][]byte) int {
	start := bytes.IndexRune(lines[0], 'S')

	splits := map[pos]struct{}{}
	positions := []pos{{start, 1}}
	for len(positions) > 0 {
		p := positions[0]
		positions = positions[1:]

		if p.y == len(lines) {
			continue
		}

		if _, ok := splits[p]; ok {
			continue
		}

		if lines[p.y][p.x] == '^' {
			splits[p] = struct{}{}
			positions = append(positions, pos{p.x - 1, p.y + 1}, pos{p.x + 1, p.y + 1})
		} else {
			positions = append(positions, pos{p.x, p.y + 1})
		}
	}

	return len(splits)
}

var seen = map[pos]int{}

func recur(lines [][]byte, x, y int) int {
	if v, ok := seen[pos{x, y}]; ok {
		return v
	}

	if y+1 == len(lines) {
		return 1
	}

	var total int

	if lines[y+1][x] == '.' {
		total = recur(lines, x, y+1)
	} else {
		total = recur(lines, x-1, y+1) + recur(lines, x+1, y+1)
	}

	seen[pos{x, y}] = total
	return total
}

func part2(lines [][]byte) int {
	return recur(lines, bytes.IndexRune(lines[0], 'S'), 0)
}

func main() {
	bs := lib.Must(os.ReadFile("inputs/day07.txt"))
	lines := bytes.Split(bs, []byte("\n"))

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}
