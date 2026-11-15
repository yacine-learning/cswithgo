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
// (Same loop as up in Go's container/heap.)
func (h MinHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h[i] >= h[parent] {
			break
		}
		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}

// down moves h[i] towards the leaves while a child is smaller, swapping with the
// smaller child. (Same loop as down in Go's container/heap.)
func (h MinHeap) down(i int) {
	for {
		smallest := i
		left, right := 2*i+1, 2*i+2
		if left < len(h) && h[left] < h[smallest] {
			smallest = left
		}
		if right < len(h) && h[right] < h[smallest] {
			smallest = right
		}
		if smallest == i {
			return
		}
		h[i], h[smallest] = h[smallest], h[i]
		i = smallest
	}
}
