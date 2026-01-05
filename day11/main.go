package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

func inputToGraph(input []byte) map[string][]string {
	out := make(map[string][]string)
	for l := range bytes.Lines(input) {
		in, outs, _ := bytes.Cut(l, []byte(":"))
		out[string(in)] = strings.Fields(string(outs))
	}
	return out
}

func dfs(graph map[string][]string, start, end string) [][]string {
	var out [][]string
	q := [][]string{{start}}

	for len(q) > 0 {
		cur := q[len(q)-1]
		q = q[:len(q)-1]
		nxt := cur[len(cur)-1]

		if nxt == end {
			out = append(out, cur)
			continue
		}

		for _, n := range graph[nxt] {
			c := make([]string, len(cur))
			copy(cur, c)
			q = append(q, append(c, n))
		}
	}

	return out
}

func part1(input []byte) int {
	graph := inputToGraph(input)

	return len(dfs(graph, "you", "out"))
}

func part2(input []byte) int {
	graph := inputToGraph(input)

	paths := dfs(graph, "svr", "out")

	var cnt int

	for _, p := range paths {
		if slices.Contains(p, "fft") && slices.Contains(p, "dac") {
			cnt++
		}
	}

	return cnt
}

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("part 1:", part1(input))
	// OOMs
	// fmt.Println("part 2:", part2(input))
}
