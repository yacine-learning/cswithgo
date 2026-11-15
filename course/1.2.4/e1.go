package main

// Ring is a queue of strings with a fixed capacity, stored in a ring buffer.
type Ring struct {
	buf   []string
	head  int // index of the front element
	count int // number of elements in the queue
}

func NewRing(capacity int) *Ring { return &Ring{buf: make([]string, capacity)} }

// Enqueue adds x at the back. It returns false if the queue is full.
func (r *Ring) Enqueue(x string) bool {
	if r.count == len(r.buf) {
		return false
	}
	back := (r.head + r.count) % len(r.buf) // wraps to the start of the array
	r.buf[back] = x
	r.count++
	return true
}

// Dequeue removes and returns the front element. ok is false if the queue is empty.
func (r *Ring) Dequeue() (x string, ok bool) {
	if r.count == 0 {
		return "", false
	}
	x = r.buf[r.head]
	r.buf[r.head] = "" // forget it, so it can be garbage collected
	r.head = (r.head + 1) % len(r.buf)
	r.count--
	return x, true
}

func (r *Ring) Len() int { return r.count }
