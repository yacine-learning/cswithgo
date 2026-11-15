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
	out := make([]Rec, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].Key <= b[j].Key {
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
