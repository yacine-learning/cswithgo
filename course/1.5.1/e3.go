package main

// FirstBad returns the first build in 0..n-1 for which isBad is true. Builds before it
// are good and builds after it are bad. It returns n when no build is bad.
// isBad is slow: call it about log2(n) times, not n times.
func FirstBad(n int, isBad func(build int) bool) int {
	return n
}
