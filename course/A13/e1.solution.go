package main

import "container/heap"

// Edge is a link to another node and what it costs to cross it (a one-way link).
type Edge struct {
	To   string
	Cost int
}

// Graph maps each node to its outgoing links.
type Graph map[string][]Edge

// ShortestPath returns the cheapest path from src to dst, its cost, and ok = false when
// dst cannot be reached. Costs are never negative.
//
// Dijkstra, 1956: always settle the closest node not settled yet. Its distance can no longer
// improve, because any other route would have to go through a node that is farther away.
func ShortestPath(g Graph, src, dst string) (path []string, cost int, ok bool) {
	dist := map[string]int{src: 0} // best distance found so far
	prev := map[string]string{}    // prev[n]: the node before n on that best path
	done := map[string]bool{}      // settled: the distance is final
	q := &queue{{src, 0}}
	for q.Len() > 0 {
		it := heap.Pop(q).(item)
		if done[it.node] {
			continue
		}
		done[it.node] = true
		if it.node == dst {
			break
		}
		for _, e := range g[it.node] {
			d := it.dist + e.Cost
			if old, seen := dist[e.To]; !seen || d < old {
				dist[e.To] = d
				prev[e.To] = it.node
				heap.Push(q, item{e.To, d})
			}
		}
	}
	if !done[dst] {
		return nil, 0, false
	}
	for n := dst; n != src; n = prev[n] {
		path = append(path, n)
	}
	path = append(path, src)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, dist[dst], true
}

// The priority queue: nodes with a tentative distance, the smallest on top (D10, container/heap).
type item struct {
	node string
	dist int
}
type queue []item

func (q queue) Len() int           { return len(q) }
func (q queue) Less(i, j int) bool { return q[i].dist < q[j].dist }
func (q queue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *queue) Push(x any)        { *q = append(*q, x.(item)) }
func (q *queue) Pop() any          { old := *q; it := old[len(old)-1]; *q = old[:len(old)-1]; return it }
