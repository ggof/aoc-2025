package lib

import "strings"

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}

	return a
}

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

func Mod(a, b int) int {
	m := a % b
	if m < 0 && b < 0 {
		m -= b
	} else if m < 0 && b > 0 {
		m += b
	}
	return m
}

func Abs[N int | uint | float32 | float64](n N) N {
	if n < 0 {
		n = 0 - n
	}
	return n
}
