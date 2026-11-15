package main

import (
	"fmt"
	"time"
)

func check(name string, xs []int, want bool) {
	got := HasDuplicate(xs)
	if got == want {
		fmt.Println("PASS", name)
	} else {
		fmt.Printf("FAIL %s: HasDuplicate(%v) returned %v, expected %v\n", name, xs, got, want)
	}
}

func main() {
	check("empty slice", []int{}, false)
	check("one value", []int{7}, false)
	check("all different", []int{3, 1, 2}, false)
	check("duplicate next to each other", []int{4, 4}, true)
	check("duplicate far apart", []int{5, 1, 2, 3, 5}, true)
	check("zero and negative values", []int{0, -1, 2, 0}, true)
	check("negative values, all different", []int{-3, -2, -1, 0}, false)

	n := 300000
	xs := make([]int, n)
	for i := range xs {
		xs[i] = i
	}
	t0 := time.Now()
	got := HasDuplicate(xs)
	d := time.Since(t0)
	switch {
	case got:
		fmt.Printf("FAIL 300 000 different values: returned true, expected false\n")
	case d > 1500*time.Millisecond:
		fmt.Printf("FAIL 300 000 different values: took %s; an O(n) check takes a few milliseconds\n", d.Round(time.Millisecond))
	default:
		fmt.Printf("PASS 300 000 different values (%s)\n", d.Round(time.Millisecond))
	}
}
