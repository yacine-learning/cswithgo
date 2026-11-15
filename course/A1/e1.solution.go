package main

// HasDuplicate reports whether a value appears twice in xs, in O(n).
func HasDuplicate(xs []int) bool {
	seen := make(map[int]bool, len(xs))
	for _, x := range xs {
		if seen[x] {
			return true
		}
		seen[x] = true
	}
	return false
}
