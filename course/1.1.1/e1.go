package main

// Steps counts the comparisons HasDuplicateSlow makes.
var Steps int

// HasDuplicateSlow reports whether a value appears twice in xs.
// It compares every pair: each element with every element after it.
func HasDuplicateSlow(xs []int) bool {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			Steps++
			if xs[i] == xs[j] {
				return true
			}
		}
	}
	return false
}
