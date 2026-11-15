package main

// Entry is one event at time T (microseconds since the program started).
type Entry struct {
	T   int
	Msg string
}

// Merge returns the entries of a and b (each sorted by T) in time order.
// On equal times, the entry from a comes first.
func Merge(a, b []Entry) []Entry {
	return nil
}
