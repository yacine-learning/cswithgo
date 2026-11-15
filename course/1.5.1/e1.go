package main

import "fmt"

// Search returns the index of target in the sorted slice a, or -1.
func Search(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right { // [left, right] is not empty: something left to look at
		mid := left + (right-left)/2 // not (left+right)/2: that can overflow
		fmt.Printf("  left=%d right=%d mid=%d a[mid]=%d\n", left, right, mid, a[mid])
		switch {
		case a[mid] == target:
			return mid
		case a[mid] < target:
			left = mid + 1 // mid and everything before it are too small
		default:
			right = mid - 1 // mid and everything after it are too big
		}
	}
	return -1
}
