package main

import "slices"

// Window is a maintenance window [Start, End], in minutes.
type Window struct{ Start, End int }

// MergeWindows returns the union of the windows as non-overlapping windows sorted by
// Start. Windows that overlap or touch (one ends exactly when the next starts) merge.
// The caller's slice is not changed.
func MergeWindows(ws []Window) []Window {
	sorted := slices.Clone(ws)
	slices.SortFunc(sorted, func(a, b Window) int { return a.Start - b.Start })
	var out []Window
	for _, w := range sorted {
		if n := len(out); n > 0 && w.Start <= out[n-1].End {
			out[n-1].End = max(out[n-1].End, w.End)
		} else {
			out = append(out, w)
		}
	}
	return out
}
