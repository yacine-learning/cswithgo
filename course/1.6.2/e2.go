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
	// hole 1: for j from lo to hi-1, swap each value smaller than the pivot into a[i], and move i

	// hole 2: put the pivot (a[hi]) at i, and return i
	_ = pivot // delete this line once you use pivot
	return i
}
