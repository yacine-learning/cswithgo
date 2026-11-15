package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rng := rand.New(rand.NewSource(42))
	h := make([]int, 100000) // heights of 100 000 trees in a forest, in cm: a few giants
	for i := range h {
		h[i] = 800 + rng.Intn(1500)
		if rng.Intn(100) == 0 {
			h[i] = 3000 + rng.Intn(3000)
		}
	}
	calls := 0
	less := func(x, y int) bool { calls++; return x < y }
	for _, p := range []int{50, 90, 99} {
		calls = 0
		k := len(h)*p/100 - 1
		fmt.Printf("p%d = %d cm   (%d comparisons for %d values)\n", p, Select(h, k, less), calls, len(h))
	}
}
