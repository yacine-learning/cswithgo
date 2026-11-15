package main

// Limiter allows at most Limit requests in any window of Window seconds.
type Limiter struct {
	Limit, Window int
	times         []int // times of the allowed requests still in the window, oldest first
}

// Allow is called for each request, with non-decreasing times (seconds).
// It returns true if the request is allowed, false if it must be refused.
func (l *Limiter) Allow(now int) bool {
	// hole 1: drop from the front the times that are no longer in the window

	// hole 2: allow (and record) the request if there is room, refuse it otherwise
	return true
}
