package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s, runs, rounds := ExternalSort([]string{"g", "c", "a", "h", "e", "b", "f", "d"}, 3)
	fmt.Println(s, "runs:", runs, "rounds:", rounds)

	rng := rand.New(rand.NewSource(1))
	big := make([]string, 100000)
	for i := range big {
		big[i] = fmt.Sprintf("2026-11-%02d %06d", 1+rng.Intn(14), rng.Intn(1000000))
	}
	for _, mem := range []int{100000, 25000, 1000, 100} {
		_, runs, rounds := ExternalSort(big, mem)
		fmt.Printf("100 000 lines, memory %6d lines: %4d runs, %d merge rounds\n", mem, runs, rounds)
	}
}
