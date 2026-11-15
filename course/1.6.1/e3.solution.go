package main

import "slices"

// Interval is a time a thread was running [Start, End], in microseconds.
type Interval struct{ Start, End int }

// MergeIntervals returns the union of the windows as non-overlapping windows sorted by
// Start. Intervals that overlap or touch (one ends exactly when the next starts) merge.
// The caller's slice is not changed.
func MergeIntervals(ws []Interval) []Interval {
	sorted := slices.Clone(ws)
	slices.SortFunc(sorted, func(a, b Interval) int { return a.Start - b.Start })
	var out []Interval
	for _, w := range sorted {
		if n := len(out); n > 0 && w.Start <= out[n-1].End {
			out[n-1].End = max(out[n-1].End, w.End)
		} else {
			out = append(out, w)
		}
	}
	return out
}
