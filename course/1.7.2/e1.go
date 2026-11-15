package main

// Entry is one log line written by a pod at time T.
type Entry struct {
	T   int
	Msg string
}

// MergeAll merges the logs of all pods into one slice sorted by T.
// Each pods[i] is already sorted by T. On equal times, the entry of the pod with the
// smaller index comes first; entries of one pod keep their order.
// Use container/heap with one entry per pod: O(N log k).
func MergeAll(pods [][]Entry) []Entry {
	return nil
}
