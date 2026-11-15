package main

import "container/heap"

// Entry is one log line written by a pod at time T.
type Entry struct {
	T   int
	Msg string
}

type cursor struct{ pod, i int }

type heads struct {
	pods [][]Entry
	cs   []cursor
}

func (h heads) Len() int { return len(h.cs) }
func (h heads) Less(a, b int) bool {
	ea, eb := h.pods[h.cs[a].pod][h.cs[a].i], h.pods[h.cs[b].pod][h.cs[b].i]
	if ea.T != eb.T {
		return ea.T < eb.T
	}
	return h.cs[a].pod < h.cs[b].pod
}
func (h heads) Swap(a, b int) { h.cs[a], h.cs[b] = h.cs[b], h.cs[a] }
func (h *heads) Push(x any)   { h.cs = append(h.cs, x.(cursor)) }
func (h *heads) Pop() any {
	x := h.cs[len(h.cs)-1]
	h.cs = h.cs[:len(h.cs)-1]
	return x
}

// MergeAll merges the logs of all pods into one slice sorted by T.
// Each pods[i] is already sorted by T. On equal times, the entry of the pod with the
// smaller index comes first; entries of one pod keep their order.
// Use container/heap with one entry per pod: O(N log k).
func MergeAll(pods [][]Entry) []Entry {
	h := &heads{pods: pods}
	total := 0
	for p, log := range pods {
		total += len(log)
		if len(log) > 0 {
			h.cs = append(h.cs, cursor{p, 0})
		}
	}
	heap.Init(h)
	out := make([]Entry, 0, total)
	for h.Len() > 0 {
		c := h.cs[0]
		out = append(out, pods[c.pod][c.i])
		if c.i+1 < len(pods[c.pod]) {
			h.cs[0].i++
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return out
}
