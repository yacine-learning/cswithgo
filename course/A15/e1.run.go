package main

import (
	"fmt"
	"strings"
)

// recursive is the same definition without a table: it recomputes the same cells again and again.
func recursive(x, y []rune, calls *int) int {
	*calls++
	if len(x) == 0 {
		return len(y)
	}
	if len(y) == 0 {
		return len(x)
	}
	change := 1
	if x[len(x)-1] == y[len(y)-1] {
		change = 0
	}
	return min(recursive(x[:len(x)-1], y, calls)+1, recursive(x, y[:len(y)-1], calls)+1,
		recursive(x[:len(x)-1], y[:len(y)-1], calls)+change)
}

func main() {
	for _, p := range [][2]string{{"kitten", "sitting"}, {"flaw", "lawn"}, {"intention", "execution"}, {"func", "fucn"}, {"", "go"}} {
		fmt.Printf("%-10s → %-10s %d\n", p[0], p[1], EditDistance(p[0], p[1]))
	}
	fmt.Println("\nthe same definition without the table:")
	for _, n := range []int{4, 6, 8, 10} {
		a, b := strings.Repeat("ab", n/2), strings.Repeat("ba", n/2)
		calls := 0
		recursive([]rune(a), []rune(b), &calls)
		fmt.Printf("  two words of %2d letters: %9d calls, against %3d cells in the table\n", n, calls, (n+1)*(n+1))
	}
}
