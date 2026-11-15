package main

// Counter reports how many clicks happened in the last Window seconds.
type Counter struct {
	Window int
	times  []int // times of the clicks still in the window, oldest first
}

// Click records a click at time now (times never decrease) and returns the number
// of clicks in the window (now-Window, now], this one included.
func (c *Counter) Click(now int) int {
	for len(c.times) > 0 && c.times[0] <= now-c.Window {
		c.times = c.times[1:]
	}
	c.times = append(c.times, now)
	return len(c.times)
}
