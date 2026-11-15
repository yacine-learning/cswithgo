package main

import (
	"math/rand"
	"slices"
)

// Select returns the k-th smallest value of a (k = 0 is the smallest), comparing values
// only with less. It does not change a. Average O(n): quickselect with a random pivot.
func Select(a []int, k int, less func(x, y int) bool) int {
	a = slices.Clone(a)
	lo, hi := 0, len(a)-1
	for lo < hi {
		r := lo + rand.Intn(hi-lo+1)
		a[r], a[hi] = a[hi], a[r]
		pivot := a[hi]
		i := lo
		for j := lo; j < hi; j++ {
			if less(a[j], pivot) {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[hi] = a[hi], a[i]
		switch {
		case k == i:
			return a[i]
		case k < i:
			hi = i - 1
		default:
			lo = i + 1
		}
	}
	return a[lo]
}
