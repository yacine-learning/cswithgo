package main

// Comparisons counts the comparisons made by merge.
var Comparisons int

// MergeSort returns a sorted copy of a.
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

// merge combines two sorted slices into one sorted slice.
func merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		Comparisons++
		if a[i] <= b[j] {
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
