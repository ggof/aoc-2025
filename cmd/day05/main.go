package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"ggof.xyz/aoc2025/lib"
)

type pair struct {
	a, b int
}

func between(p pair, i int) bool {
	return p.a <= i && p.b >= i
}

func overlaps(a, b pair) (pair, bool) {
	if a.b >= b.a && a.a <= b.b {
		// ovelap exists

	}

	return pair{}, false
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

func part2(input []byte) int {
	rangesStr, _, _ := bytes.Cut(input, []byte("\n\n"))
	rangesParts := bytes.Split(bytes.TrimSuffix(rangesStr, []byte("\n")), []byte("\n"))

	ranges := make([]pair, len(rangesParts))
	for i, rp := range rangesParts {
		fmt.Sscanf(string(rp), "%d-%d", &ranges[i].a, &ranges[i].b)
	}

	var rangesWithoutOverlap []pair
	for _, r := range ranges {

		for _, rwo := range rangesWithoutOverlap {

		}
		fmt.Println("range has len", r.b - r.a)
	}

	var total int
	return total 
}

func main() {
	file := lib.Must(os.Open("inputs/day05.txt"))
	bs := lib.Must(io.ReadAll(file))
	fmt.Println("part 1:", part1(bs))
	fmt.Println("part 2:", part2(bs))

}
