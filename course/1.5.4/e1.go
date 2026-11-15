package main

// All three functions receive ts: detection times in seconds, sorted, possibly repeated.

// Count returns how many timestamps t satisfy from <= t < to.
// It is called many times on the same ts: make it O(log n).
func Count(ts []int, from, to int) int {
	return 0
}

// Busiest returns the window [start, start+width) with the most timestamps, where start
// is one of the timestamps, and how many it holds. On a tie, the earliest start.
// An empty ts gives 0, 0.
func Busiest(ts []int, width int) (start, count int) {
	return 0, 0
}

// Silences returns every gap of at least min seconds between two consecutive
// timestamps, as [from, to] pairs, in order.
func Silences(ts []int, min int) [][2]int {
	return nil
}
