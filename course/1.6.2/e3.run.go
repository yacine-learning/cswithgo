package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rng := rand.New(rand.NewSource(42))
	lat := make([]int, 100000) // request latencies in ms: mostly fast, a slow tail
	for i := range lat {
		lat[i] = 5 + rng.Intn(20)
		if rng.Intn(100) == 0 {
			lat[i] = 200 + rng.Intn(800)
		}
	}
	calls := 0
	less := func(x, y int) bool { calls++; return x < y }
	for _, p := range []int{50, 90, 99} {
		calls = 0
		k := len(lat)*p/100 - 1
		fmt.Printf("p%d = %d ms   (%d comparisons for %d values)\n", p, Select(lat, k, less), calls, len(lat))
	}
}
