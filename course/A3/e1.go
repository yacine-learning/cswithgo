package main

// Search returns the index of target in the sorted slice a, or -1.
func Search(a []int, target int) int {
	left, right := 0, len(a)-1
	for left > right { // hole: loop while [left, right] still holds something
		mid := left + (right-left)/2
		switch {
		case a[mid] == target:
			return mid
		case a[mid] < target:
			left = left // hole: mid is too small, and so is everything before it
		default:
			right = right // hole: mid is too big, and so is everything after it
		}
	}
	return -1
}
