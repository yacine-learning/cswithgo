package main

import "container/heap"

// Entry is one reading made by a station at time T.
type Entry struct {
	T   int
	Msg string
}

type cursor struct{ station, i int }

type heads struct {
	stations [][]Entry
	cs   []cursor
}

func (h heads) Len() int { return len(h.cs) }
func (h heads) Less(a, b int) bool {
	ea, eb := h.stations[h.cs[a].station][h.cs[a].i], h.stations[h.cs[b].station][h.cs[b].i]
	if ea.T != eb.T {
		return ea.T < eb.T
	}
	return h.cs[a].station < h.cs[b].station
}
func (h heads) Swap(a, b int) { h.cs[a], h.cs[b] = h.cs[b], h.cs[a] }
func (h *heads) Push(x any)   { h.cs = append(h.cs, x.(cursor)) }
func (h *heads) Pop() any {
	x := h.cs[len(h.cs)-1]
	h.cs = h.cs[:len(h.cs)-1]
	return x
}

// MergeAll merges the records of all stations into one slice sorted by T.
// Each stations[i] is already sorted by T. On equal times, the entry of the station with
// the smaller index comes first; entries of one station keep their order.
// Use container/heap with one entry per station: O(N log k).
func MergeAll(stations [][]Entry) []Entry {
	h := &heads{stations: stations}
	total := 0
	for p, log := range stations {
		total += len(log)
		if len(log) > 0 {
			h.cs = append(h.cs, cursor{p, 0})
		}
	}
	heap.Init(h)
	out := make([]Entry, 0, total)
	for h.Len() > 0 {
		c := h.cs[0]
		out = append(out, stations[c.station][c.i])
		if c.i+1 < len(stations[c.station]) {
			h.cs[0].i++
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return out
}
