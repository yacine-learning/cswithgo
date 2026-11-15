package main

import (
	"container/heap"
	"slices"
)

type item struct {
	msg   string
	count int
}

// minH is a min-heap where the top is the item to drop first:
// the smallest count, and on equal counts the name that sorts last.
type minH []item

func (h minH) Len() int { return len(h) }
func (h minH) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	}
	return h[i].msg > h[j].msg
}
func (h minH) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minH) Push(x any)   { *h = append(*h, x.(item)) }
func (h *minH) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TopK returns the k messages with the highest counts, most frequent first;
// equal counts in alphabetical order. Use container/heap with at most k entries.
func TopK(counts map[string]int, k int) []string {
	if k <= 0 {
		return nil
	}
	h := &minH{}
	for msg, c := range counts {
		heap.Push(h, item{msg, c})
		if h.Len() > k {
			heap.Pop(h) // drop the smallest: it cannot be in the top k
		}
	}
	var out []string
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(item).msg)
	}
	slices.Reverse(out)
	return out
}
