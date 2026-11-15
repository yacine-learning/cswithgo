package main

// Entry is one reading made by a station at time T.
type Entry struct {
	T   int
	Msg string
}

// MergeAll merges the records of all stations into one slice sorted by T.
// Each stations[i] is already sorted by T. On equal times, the entry of the station with
// the smaller index comes first; entries of one station keep their order.
// Use container/heap with one entry per station: O(N log k).
func MergeAll(stations [][]Entry) []Entry {
	return nil
}
