package main

// Limiter allows at most Limit requests in any window of Window seconds.
type Limiter struct {
	Limit, Window int
	times         []int // times of the allowed requests still in the window, oldest first
}

// Allow is called for each request, with non-decreasing times (seconds).
// It returns true if the request is allowed, false if it must be refused.
func (l *Limiter) Allow(now int) bool {
	for len(l.times) > 0 && l.times[0] <= now-l.Window {
		l.times = l.times[1:]
	}
	if len(l.times) < l.Limit {
		l.times = append(l.times, now)
		return true
	}
	return false
}
