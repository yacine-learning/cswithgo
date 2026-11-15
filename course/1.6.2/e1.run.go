package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	a := []int{38, 27, 43, 3, 9, 82, 10}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)

	n := 4000
	random := rand.New(rand.NewSource(1)).Perm(n)
	sorted := make([]int, n)
	for i := range sorted {
		sorted[i] = i
	}
	try := func(name string, in []int, rp bool) {
		b := append([]int(nil), in...)
		Comparisons, RandomPivot = 0, rp
		QuickSort(b, 0, n-1)
		fmt.Printf("%-34s %9d comparisons\n", name, Comparisons)
	}
	fmt.Printf("\nn = %d   (n·log2 n ≈ %.0f, n²/2 = %d)\n", n, float64(n)*math.Log2(float64(n)), n*n/2)
	try("random input, last element pivot", random, false)
	try("sorted input, last element pivot", sorted, false)
	try("sorted input, random pivot", sorted, true)
}
