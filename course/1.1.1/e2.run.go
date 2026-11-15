package main

import (
	"fmt"
	"time"
)

func slow(xs []int) bool {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[i] == xs[j] {
				return true
			}
		}
	}
	return false
}

func distinct(n int) []int {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = i
	}
	return xs
}

func main() {
	fmt.Println("HasDuplicate([3 1 4 1 5]) =", HasDuplicate([]int{3, 1, 4, 1, 5}))
	fmt.Println("HasDuplicate([2 7 1 8])   =", HasDuplicate([]int{2, 7, 1, 8}))
	fmt.Println()
	fmt.Println("n          slow      fast")
	for n := 4000; n <= 32000; n *= 2 {
		xs := distinct(n)
		t0 := time.Now()
		slow(xs)
		ds := time.Since(t0)
		t0 = time.Now()
		HasDuplicate(xs)
		df := time.Since(t0)
		fmt.Printf("%-10d %-9s %s\n", n, ds.Round(time.Millisecond), df.Round(10*time.Microsecond))
	}
	fmt.Println("\nfast only, much bigger n:")
	for n := 250000; n <= 2000000; n *= 2 {
		xs := distinct(n)
		t0 := time.Now()
		HasDuplicate(xs)
		fmt.Printf("%-10d %s\n", n, time.Since(t0).Round(time.Millisecond))
	}
}
