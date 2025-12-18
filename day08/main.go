package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"ggof.xyz/aoc2025/lib"
)

type p3 struct{ x, y, z int }
type pdist struct {
	fromIndex int
	toIndex   int
	dist      int
}

type ByDistance []pdist

func (bd ByDistance) Len() int           { return len(bd) }
func (bd ByDistance) Swap(i, j int)      { bd[j], bd[i] = bd[i], bd[j] }
func (bd ByDistance) Less(i, j int) bool { return bd[i].dist < bd[j].dist }

type BySize [][]int

func (bs BySize) Len() int           { return len(bs) }
func (bs BySize) Swap(i, j int)      { bs[j], bs[i] = bs[i], bs[j] }
func (bs BySize) Less(i, j int) bool { return len(bs[i]) > len(bs[j]) }

func dist(p1, p2 p3) int {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dz := p2.z - p1.z

	return (dx * dx) + (dy * dy) + (dz * dz)
}

func part1(lines []string, take int) int {
	ps := make([]p3, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, ",")
		ps[i].x = lib.Must(strconv.Atoi(parts[0]))
		ps[i].y = lib.Must(strconv.Atoi(parts[1]))
		ps[i].z = lib.Must(strconv.Atoi(parts[2]))
	}

	dists := make([]pdist, 0)
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			p2 := ps[j]
			dists = append(dists, pdist{i, j, dist(p1, p2)})
		}
	}

	sort.Sort(ByDistance(dists))
	dists = dists[:take]

	// make an array of the indexes of the points
	circuits := make([][]int, 0)
	for _, d := range dists {
		circuitForFromIndex := slices.IndexFunc(circuits, func(c []int) bool { return slices.Contains(c, d.fromIndex) })
		circuitForToIndex := slices.IndexFunc(circuits, func(c []int) bool { return slices.Contains(c, d.toIndex) })

		switch {
		case circuitForFromIndex != -1 && circuitForToIndex != -1 && circuitForFromIndex != circuitForToIndex:
			// concat them
			circuits[circuitForFromIndex] = append(circuits[circuitForFromIndex], circuits[circuitForToIndex]...)

			// swap with last + remove
			circuits[circuitForToIndex] = circuits[0]
			circuits = circuits[1:]
		case circuitForFromIndex == -1 && circuitForToIndex != -1:
			circuits[circuitForToIndex] = append(circuits[circuitForToIndex], d.fromIndex)
		case circuitForFromIndex != -1 && circuitForToIndex == -1:
			circuits[circuitForFromIndex] = append(circuits[circuitForFromIndex], d.toIndex)
		case circuitForFromIndex == -1 && circuitForToIndex == -1:
			circuits = append(circuits, []int{d.fromIndex, d.toIndex})
		}
	}

	sort.Sort(BySize(circuits))

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(lines []string) int {

	ps := make([]p3, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, ",")
		ps[i].x = lib.Must(strconv.Atoi(parts[0]))
		ps[i].y = lib.Must(strconv.Atoi(parts[1]))
		ps[i].z = lib.Must(strconv.Atoi(parts[2]))
	}

	dists := make([]pdist, 0)
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			p2 := ps[j]
			dists = append(dists, pdist{i, j, dist(p1, p2)})
		}
	}

	sort.Sort(ByDistance(dists))

	// make an array of the indexes of the points
	circuits := make([][]int, 0)
	for _, d := range dists {
		circuitForFromIndex := slices.IndexFunc(circuits, func(c []int) bool { return slices.Contains(c, d.fromIndex) })
		circuitForToIndex := slices.IndexFunc(circuits, func(c []int) bool { return slices.Contains(c, d.toIndex) })

		switch {
		case circuitForFromIndex != -1 && circuitForToIndex != -1 && circuitForFromIndex != circuitForToIndex:
			// concat them
			circuits[circuitForFromIndex] = append(circuits[circuitForFromIndex], circuits[circuitForToIndex]...)

			// swap with last + remove
			circuits[circuitForToIndex] = circuits[0]
			circuits = circuits[1:]
		case circuitForFromIndex == -1 && circuitForToIndex != -1:
			circuits[circuitForToIndex] = append(circuits[circuitForToIndex], d.fromIndex)
		case circuitForFromIndex != -1 && circuitForToIndex == -1:
			circuits[circuitForFromIndex] = append(circuits[circuitForFromIndex], d.toIndex)
		case circuitForFromIndex == -1 && circuitForToIndex == -1:
			circuits = append(circuits, []int{d.fromIndex, d.toIndex})
		}

		if len(circuits) == 1 && len(circuits[0])	== len(lines) {
			return ps[d.fromIndex].x * ps[d.toIndex].x
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
