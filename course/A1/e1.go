package main

// HasDuplicate reports whether a value appears twice in xs, in O(n).
// It remembers every value it has seen in a map, so checking a new value
// is one step instead of a comparison with all the others.
func HasDuplicate(xs []int) bool {
	seen := make(map[int]bool, len(xs))
	// Go refuses to compile a variable that is never used. This line "uses" seen
	// by throwing the value away, so the file compiles before you fill the holes.
	// Delete it once your code uses seen.
	_ = seen
	for _, x := range xs {
		// Go refuses to compile a variable that is never used. This line "uses" x
		// by throwing the value away, so the file compiles before you fill the holes.
		// Delete it once your code uses x.
		_ = x

		// hole 1: if x was seen already, there is a duplicate

		// hole 2: remember x

	}
	return false
}
