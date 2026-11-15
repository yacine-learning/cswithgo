package main

// Search returns the index of target in the sorted slice a, or -1.
func Search(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case a[mid] == target:
			return mid
		case a[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}
