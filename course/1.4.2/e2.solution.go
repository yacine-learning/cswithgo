package main

// Graph maps each node to its neighbours.
type Graph map[string][]string

// ShortestPath returns the path from -> to with the fewest edges (from and to included),
// or nil if to cannot be reached.
func ShortestPath(g Graph, from, to string) []string {
	parent := map[string]string{}
	seen := map[string]bool{from: true}
	queue := []string{from}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range g[cur] {
			if !seen[next] {
				seen[next] = true
				parent[next] = cur
				queue = append(queue, next)
			}
		}
	}
	if !seen[to] {
		return nil
	}
	path := []string{to}
	for n := to; n != from; {
		n = parent[n]
		path = append(path, n)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
