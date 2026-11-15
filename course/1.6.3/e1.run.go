package main

import (
	"fmt"
	"math/rand"
)

// word makes a random pronounceable word, like an entry of an imaginary dictionary.
func word(rng *rand.Rand) string {
	c, v := "bcdfghklmnprstvz", "aeiou"
	w := []byte{}
	for i := 0; i < 2+rng.Intn(3); i++ {
		w = append(w, c[rng.Intn(len(c))], v[rng.Intn(len(v))])
	}
	return string(w)
}

func main() {
	s, runs, rounds := ExternalSort([]string{"g", "c", "a", "h", "e", "b", "f", "d"}, 3)
	fmt.Println(s, "runs:", runs, "rounds:", rounds)

	rng := rand.New(rand.NewSource(1))
	big := make([]string, 100000)
	for i := range big {
		big[i] = word(rng)
	}
	for _, mem := range []int{100000, 25000, 1000, 100} {
		_, runs, rounds := ExternalSort(big, mem)
		fmt.Printf("100 000 words, memory %6d words: %4d runs, %d merge rounds\n", mem, runs, rounds)
	}
}
