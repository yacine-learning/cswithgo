package main

// Comparisons counts the comparisons made by partition.
var Comparisons int

// QuickSort sorts a[lo..hi] in place.
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}

// partition puts a pivot at its final place p: a[lo..p-1] < pivot <= a[p+1..hi]. (Lomuto)
func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo // a[lo..i-1] holds the values smaller than the pivot
	for j := lo; j < hi; j++ {
		Comparisons++
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i] // the pivot goes just after the smaller ones
	return i
}
