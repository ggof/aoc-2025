package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"ggof.xyz/aoc2025/lib"
)

type p2 struct{ x, y int }
type parea struct {
	p1   p2
	p2   p2
	area int
}

func area(p1, p2 p2) int {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	return (lib.Abs(dx) + 1) * (lib.Abs(dy) + 1)
}

func part1(lines []string) int {
	ps := make([]p2, len(lines))
	for i, l := range lines {
		fmt.Sscanf(l, "%d,%d,%d", &ps[i].x, &ps[i].y)
	}

	areas := make([]parea, 0, len(ps)*len(ps))
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			p2 := ps[j]
			areas = append(areas, parea{p1, p2, area(p1, p2)})
		}
	}

	return slices.MaxFunc(areas, func(a, b parea) int { return a.area - b.area }).area
}

func main() {
	bs := lib.Must(os.ReadFile("inputs/day09.txt"))
	lines := strings.Split(string(bs), "\n")
	fmt.Println("part 1:", part1(lines))
}
