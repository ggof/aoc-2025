package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"ggof.xyz/aoc2025/lib"
)

type p3 struct{ x, y, z int }
type pdist struct {
	p1   p3
	p2   p3
	dist int
}

type byDistance []pdist

func (bd byDistance) Len() int           { return len(bd) }
func (bd byDistance) Swap(i, j int)      { bd[j], bd[i] = bd[i], bd[j] }
func (bd byDistance) Less(i, j int) bool { return bd[i].dist < bd[j].dist }

type bySize [][]p3

func (bs bySize) Len() int           { return len(bs) }
func (bs bySize) Swap(i, j int)      { bs[j], bs[i] = bs[i], bs[j] }
func (bs bySize) Less(i, j int) bool { return len(bs[i]) > len(bs[j]) }

func dist(p1, p2 p3) int {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dz := p2.z - p1.z

	return (dx * dx) + (dy * dy) + (dz * dz)
}

func sortedDistances(lines []string) []pdist {
	ps := make([]p3, len(lines))
	for i, l := range lines {
		fmt.Sscanf(l, "%d,%d,%d", &ps[i].x, &ps[i].y, &ps[i].z)
	}

	dists := make([]pdist, 0, len(ps)*len(ps))
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			p2 := ps[j]
			dists = append(dists, pdist{p1, p2, dist(p1, p2)})
		}
	}

	sort.Sort(byDistance(dists))
	return dists
}

func updateCircuits(circuits [][]p3, d pdist) [][]p3 {
	p1i := slices.IndexFunc(circuits, func(c []p3) bool { return slices.Contains(c, d.p1) })
	p2i := slices.IndexFunc(circuits, func(c []p3) bool { return slices.Contains(c, d.p2) })

	switch {
	case p1i != -1 && p2i != -1 && p1i != p2i:
		circuits[p1i] = append(circuits[p1i], circuits[p2i]...)
		circuits[p2i] = circuits[0]
		circuits = circuits[1:]
	case p1i == -1 && p2i != -1:
		circuits[p2i] = append(circuits[p2i], d.p1)
	case p1i != -1 && p2i == -1:
		circuits[p1i] = append(circuits[p1i], d.p2)
	case p1i == -1 && p2i == -1:
		circuits = append(circuits, []p3{d.p1, d.p2})
	}

	return circuits
}

func part1(lines []string, take int) int {
	dists := sortedDistances(lines)
	// make an array of the indexes of the points
	var circuits [][]p3
	for _, d := range dists[:take] {
		circuits = updateCircuits(circuits, d)
	}

	sort.Sort(bySize(circuits))

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(lines []string) int {
	dists := sortedDistances(lines)
	var circuits [][]p3
	for _, d := range dists {
		circuits = updateCircuits(circuits, d)

		if len(circuits) == 1 && len(circuits[0]) == len(lines) {
			return d.p1.x * d.p2.x
		}
	}

	panic("arrived at the end without connecting all boxes")
}

func main() {
	input := lib.Must(os.ReadFile("inputs/day08.txt"))
	lines := strings.Split(string(input), "\n")

	fmt.Println("part 1:", part1(lines, 1000))
	fmt.Println("part 2:", part2(lines))
}
