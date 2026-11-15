package main

// Interval is an observation [Start, End], in minutes after dusk.
type Interval struct{ Start, End int }

// MergeIntervals returns the union of the windows as non-overlapping windows sorted by
// Start. Intervals that overlap or touch (one ends exactly when the next starts) merge.
// The caller's slice is not changed.
func MergeIntervals(ws []Interval) []Interval {
	return nil
}
