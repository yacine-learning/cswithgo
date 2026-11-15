package main

// Counter reports how many clicks happened in the last Window seconds.
type Counter struct {
	Window int
	times  []int // times of the clicks still in the window, oldest first
}

// Click records a click at time now (times never decrease) and returns the number
// of clicks in the window (now-Window, now], this one included.
func (c *Counter) Click(now int) int {
	// hole 1: drop from the front the times that are no longer in the window

	// hole 2: record now and return how many clicks are in the window
	return 0
}
