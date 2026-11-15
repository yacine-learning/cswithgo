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
		_ = it // Go refuses unused variables: delete this line once your code uses it

		// hole 1: the queue may hold old, longer entries for a node already settled: skip them.
		//         Otherwise settle it.node, and stop early if it is dst.

		// hole 2: for each edge out of it.node, if going through it.node is shorter than the
		//         best distance known for e.To (or e.To has none yet), record the new distance,
		//         its prev, and push it on the queue.
	}
	_, _, _ = dist, prev, done // delete this line once your code uses dist, prev and done

	// hole 3: if dst was never settled, it is unreachable. Otherwise walk prev back from dst
	//         to src, then reverse, to get the path.
	return nil, 0, false
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
