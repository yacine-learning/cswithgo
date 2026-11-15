package main

// MinHeap is a binary min-heap of ints in a slice: h[0] is the smallest,
// and the children of h[i] are h[2*i+1] and h[2*i+2].
type MinHeap []int

func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

// Pop removes and returns the smallest value. The heap must not be empty.
func (h *MinHeap) Pop() int {
	old := *h
	min := old[0]
	last := len(old) - 1
	old[0] = old[last] // the last value takes the root's place...
	*h = old[:last]
	h.down(0) // ...and sinks to where it belongs
	return min
}

// up moves h[i] towards the root while it is smaller than its parent.
func (h MinHeap) up(i int) {
	// hole
}

// down moves h[i] towards the leaves while a child is smaller, swapping with the
// smaller child.
func (h MinHeap) down(i int) {
	// hole
}
