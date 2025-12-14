package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"ggof.xyz/aoc2025/lib"
)

type pair struct {
	a, b int
}

func between(p pair, i int) bool {
	return p.a <= i && p.b >= i
}

func part1(input []byte) int {
	rangesStr, ingredientsStr, _ := bytes.Cut(input, []byte("\n\n"))
	rangesParts := bytes.Split(bytes.TrimSuffix(rangesStr, []byte("\n")), []byte("\n"))

	ranges := make([]pair, len(rangesParts))
	for i, rp := range rangesParts {
		fmt.Sscanf(string(rp), "%d-%d", &ranges[i].a, &ranges[i].b)
	}

	ingredientsParts := bytes.Split(bytes.TrimSuffix(ingredientsStr, []byte("\n")), []byte("\n"))
	ingredients := make([]int, len(ingredientsParts))
	for i, ip := range ingredientsParts {
		ingredients[i] = lib.Must(strconv.Atoi(string(ip)))
	}

	var total int
	for _, i := range ingredients {
		for _, p := range ranges {
			if between(p, i) {
				total++
				break
			}
		}
	}

	return total
}

type ByLen []pair

func (b ByLen) Len() int               { return len(b) }
func (b ByLen) Less(i int, j int) bool { return (b[i].b - b[i].a) > (b[j].b - b[j].a) }
func (b ByLen) Swap(i int, j int)      { b[i], b[j] = b[j], b[i] }

func part2(input []byte) int {
	rangesStr, _, _ := bytes.Cut(input, []byte("\n\n"))
	rangesParts := bytes.Split(bytes.TrimSuffix(rangesStr, []byte("\n")), []byte("\n"))

	ranges := make([]pair, len(rangesParts))
	for i, rp := range rangesParts {
		fmt.Sscanf(string(rp), "%d-%d", &ranges[i].a, &ranges[i].b)
	}

	// sort first by size of range to ensure the biggest ranges get treated first.
	// Otherwise we'd have to split them when they get here and that's no fun.
	sort.Sort(ByLen(ranges))

	var total int
	var rangesWithoutOverlap []pair
	for _, r := range ranges {
		for _, rwo := range rangesWithoutOverlap {
			if between(rwo, r.a) {
				r.a = min(rwo.b, r.b) + 1
			}

			if between(rwo, r.b) {
				r.b = max(rwo.a, r.a) - 1
			}
		}

		if r.a <= r.b {
			total += r.b - r.a + 1
			rangesWithoutOverlap = append(rangesWithoutOverlap, r)
		}
	}

	return total
}

func main() {
	file := lib.Must(os.Open("inputs/day05.txt"))
	bs := lib.Must(io.ReadAll(file))
	fmt.Println("part 1:", part1(bs))
	fmt.Println("part 2:", part2(bs))
}
