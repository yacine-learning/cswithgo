package main

import "sort"

// All three functions receive ts: error timestamps in seconds, sorted, possibly repeated.

// Count returns how many timestamps t satisfy from <= t < to.
// It is called many times on the same ts: make it O(log n).
func Count(ts []int, from, to int) int {
	if to <= from {
		return 0
	}
	return sort.SearchInts(ts, to) - sort.SearchInts(ts, from)
}

// Busiest returns the window [start, start+width) with the most timestamps, where start
// is one of the timestamps, and how many it holds. On a tie, the earliest start.
// An empty ts gives 0, 0.
func Busiest(ts []int, width int) (start, count int) {
	right := 0
	for left := range ts {
		if right < left {
			right = left
		}
		for right < len(ts) && ts[right] < ts[left]+width {
			right++
		}
		if right-left > count {
			start, count = ts[left], right-left
		}
	}
	return start, count
}

// Silences returns every gap of at least min seconds between two consecutive
// timestamps, as [from, to] pairs, in order.
func Silences(ts []int, min int) [][2]int {
	var out [][2]int
	for i := 0; i+1 < len(ts); i++ {
		if ts[i+1]-ts[i] >= min {
			out = append(out, [2]int{ts[i], ts[i+1]})
		}
	}
	return out
}
