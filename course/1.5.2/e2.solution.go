package main

// Dedup removes repeated values from the sorted slice a, in place: afterwards the
// distinct values are a[:k], in order. It returns k.
func Dedup(a []int) int {
	if len(a) == 0 {
		return 0
	}
	k := 1 // a[:k] holds the values kept so far
	for i := 1; i < len(a); i++ {
		if a[i] != a[k-1] {
			a[k] = a[i]
			k++
		}
	}
	return k
}
