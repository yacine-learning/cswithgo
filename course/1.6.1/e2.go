package main

// Rec is a record sorted by Key; ID shows the original order of equal keys.
type Rec struct {
	Key int
	ID  string
}

// MergeSort returns a copy of a sorted by Key, stable.
func MergeSort(a []Rec) []Rec {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

// merge combines two slices sorted by Key. On equal keys, a's records come first.
func merge(a, b []Rec) []Rec {
	// hole
	return nil
}
