package main

// Entry is one event recorded by a thread at time T.
type Entry struct {
	T   int
	Msg string
}

// MergeAll merges the records of all threads into one slice sorted by T.
// Each threads[i] is already sorted by T. On equal times, the entry of the thread with
// the smaller index comes first; entries of one thread keep their order.
// Use container/heap with one entry per thread: O(N log k).
func MergeAll(threads [][]Entry) []Entry {
	return nil
}
