package main

// Graph maps each node to its neighbours.
type Graph map[string][]string

// ShortestPath returns the path from -> to with the fewest edges (from and to included),
// or nil if to cannot be reached.
func ShortestPath(g Graph, from, to string) []string {
	parent := map[string]string{}
	seen := map[string]bool{from: true}
	queue := []string{from}
	// hole 1: BFS: take the front of the queue; for each neighbour not seen yet,
	// mark it seen, record its parent, and queue it
	_, _ = parent, queue // delete this line once you use parent and queue

	if !seen[to] {
		return nil
	}
	// hole 2: walk back from to through parent until from, then reverse
	return nil
}
