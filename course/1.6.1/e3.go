package main

// Window is a maintenance window [Start, End], in minutes.
type Window struct{ Start, End int }

// MergeWindows returns the union of the windows as non-overlapping windows sorted by
// Start. Windows that overlap or touch (one ends exactly when the next starts) merge.
// The caller's slice is not changed.
func MergeWindows(ws []Window) []Window {
	return nil
}
