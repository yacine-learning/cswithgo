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
	// hole: refuse when full; otherwise put x in the back slot (wrapping around) and count it
	return false
}

// Dequeue removes and returns the front element. ok is false if the queue is empty.
func (r *Ring) Dequeue() (x string, ok bool) {
	// hole: refuse when empty; otherwise take buf[head], clear that slot, move head (wrapping around)
	return "", false
}

func (r *Ring) Len() int { return r.count }
