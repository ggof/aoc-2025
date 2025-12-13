package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"ggof.xyz/aoc2025/lib"
)

func parseLine(line string) int {
	switch line[0] {
	case 'L':
		line = "-" + line[1:]
	case 'R':
		line = "+" + line[1:]
	default:
		panic("invalid line " + line)
	}
	return int(lib.Must(strconv.ParseInt(line, 10, 64)))
}

func part1(lines []string) int {
	acc := 50
	cnt := 0

	for _, l := range lines {
		amount := parseLine(l)
		acc = lib.Mod(acc+amount, 100)

		if acc == 0 {
			cnt++
		}
	}

	return cnt
}

func part2(lines []string) int {
	acc := 50
	cnt := 0

	for _, l := range lines {
		dist := parseLine(l)

		fullTurns := lib.Abs(dist / 100)
		cnt += fullTurns
		dist %= 100

		acc += dist
		if acc > 100 || (acc < 0 && acc != dist) {
			cnt++
		}

		acc = lib.Mod(acc, 100)
		if acc == 0 {
			cnt++
		}
	}

	return cnt
}

func main() {
	file := lib.Must(os.Open("inputs/day01.txt"))
	input := lib.Must(io.ReadAll(file))
	lines := lib.Lines(strings.Trim(string(input), "\n"))
	fmt.Printf("day01/part1: %d\n", part1(lines))
	fmt.Printf("day01/part2: %d\n", part2(lines))
}
