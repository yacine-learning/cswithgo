package main

import "container/heap"

// Entry is one event recorded by a thread at time T.
type Entry struct {
	T   int
	Msg string
}

type cursor struct{ thread, i int }

type heads struct {
	threads [][]Entry
	cs      []cursor
}

func (h heads) Len() int { return len(h.cs) }
func (h heads) Less(a, b int) bool {
	ea, eb := h.threads[h.cs[a].thread][h.cs[a].i], h.threads[h.cs[b].thread][h.cs[b].i]
	if ea.T != eb.T {
		return ea.T < eb.T
	}
	return h.cs[a].thread < h.cs[b].thread
}
func (h heads) Swap(a, b int) { h.cs[a], h.cs[b] = h.cs[b], h.cs[a] }
func (h *heads) Push(x any)   { h.cs = append(h.cs, x.(cursor)) }
func (h *heads) Pop() any {
	x := h.cs[len(h.cs)-1]
	h.cs = h.cs[:len(h.cs)-1]
	return x
}

// MergeAll merges the records of all threads into one slice sorted by T.
// Each threads[i] is already sorted by T. On equal times, the entry of the thread with
// the smaller index comes first; entries of one thread keep their order.
// Use container/heap with one entry per thread: O(N log k).
func MergeAll(threads [][]Entry) []Entry {
	h := &heads{threads: threads}
	total := 0
	for p, log := range threads {
		total += len(log)
		if len(log) > 0 {
			h.cs = append(h.cs, cursor{p, 0})
		}
	}
	heap.Init(h)
	out := make([]Entry, 0, total)
	for h.Len() > 0 {
		c := h.cs[0]
		out = append(out, threads[c.thread][c.i])
		if c.i+1 < len(threads[c.thread]) {
			h.cs[0].i++
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return out
}
