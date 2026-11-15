package main

import (
	"fmt"
	"math"
)

// digits of n! = floor(log10(1) + log10(2) + … + log10(n)) + 1
func factorialDigits(n int) int {
	s := 0.0
	for k := 2; k <= n; k++ {
		s += math.Log10(float64(k))
	}
	return int(s) + 1
}

func main() {
	tests := 0
	holds := func(n int) bool {
		tests++
		d := factorialDigits(n)
		fmt.Printf("  %5d! has %5d digits\n", n, d)
		return d >= 1000
	}
	n := FirstTrue(10000, holds)
	fmt.Printf("the first n with n! of at least 1 000 digits: %d, found with %d tests out of 10 000 candidates\n", n, tests)
}
