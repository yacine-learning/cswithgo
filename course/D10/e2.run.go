package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Euclid's algorithm and binary search, as Go source: which names are used most?
	text := `func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func search(xs []int, x int) int {
	left, right := 0, len(xs)
	for left < right {
		mid := left + (right-left)/2
		if xs[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}`
	counts := map[string]int{}
	for _, w := range strings.FieldsFunc(strings.ToLower(text), func(r rune) bool { return !unicode.IsLetter(r) }) {
		counts[w]++
	}
	fmt.Println("top 5:", TopK(counts, 5))
}
