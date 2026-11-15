package main

// Extend returns the path prefix followed by step.
// prefix must not change, and each result must be independent of the others.
func Extend(prefix []int, step int) []int {
	out := make([]int, len(prefix), len(prefix)+1)
	copy(out, prefix)
	return append(out, step)
}
