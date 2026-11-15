package main

// FirstTrue returns the first k in 0..n-1 for which holds(k) is true. The property is
// false before that k and true from it on. It returns n when it holds nowhere below n.
// holds is expensive: call it about log2(n) times, not n times.
func FirstTrue(n int, holds func(k int) bool) int {
	lo, hi := 0, n // the answer is in [lo, hi]
	for lo < hi {
		mid := lo + (hi-lo)/2
		if holds(mid) {
			hi = mid // true at mid: the first one is mid or before
		} else {
			lo = mid + 1 // false at mid: the first one is after
		}
	}
	return lo
}
