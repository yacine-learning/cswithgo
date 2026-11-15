package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(MergeSort([]int{38, 27, 43, 3, 9, 82, 10}))
	rng := rand.New(rand.NewSource(1))
	fmt.Println("\nn          comparisons   n·log2(n)")
	for n := 1000; n <= 1_000_000; n *= 10 {
		Comparisons = 0
		MergeSort(rng.Perm(n))
		fmt.Printf("%-10d %-13d %.0f\n", n, Comparisons, float64(n)*math.Log2(float64(n)))
	}
}
