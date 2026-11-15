package main

// Entry is one event at time T (microseconds since the program started).
type Entry struct {
	T   int
	Msg string
}

// Merge returns the entries of a and b (each sorted by T) in time order.
// On equal times, the entry from a comes first.
func Merge(a, b []Entry) []Entry {
	out := make([]Entry, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].T <= b[j].T {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	return append(out, b[j:]...)
}
