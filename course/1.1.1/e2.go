package main

// HasDuplicate reports whether a value appears twice in xs, in O(n).
// It remembers every value it has seen in a map, so checking a new value
// is one step instead of a comparison with all the others.
func HasDuplicate(xs []int) bool {
	seen := make(map[int]bool, len(xs))
	_ = seen // delete this line once you use seen
	for _, x := range xs {
		_ = x // delete this line once you use x

		// hole 1: if x was seen already, there is a duplicate

		// hole 2: remember x

	}
	return false
}
