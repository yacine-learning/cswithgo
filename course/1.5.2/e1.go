package main

import "fmt"

// PairSum returns the indices of two values of the sorted slice a that add up to target.
func PairSum(a []int, target int) (i, j int, found bool) {
	left, right := 0, len(a)-1
	for left < right {
		sum := a[left] + a[right]
		fmt.Printf("  a[%d]+a[%d] = %d+%d = %d\n", left, right, a[left], a[right], sum)
		switch {
		case sum == target:
			return left, right, true
		case sum < target:
			left++ // a[left] is too small even with the biggest partner: drop it
		default:
			right-- // a[right] is too big even with the smallest partner: drop it
		}
	}
	return 0, 0, false
}
