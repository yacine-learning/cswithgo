package main

// Graph maps each node to its neighbours.
type Graph map[string][]string

// BFS visits the nodes reachable from start, nearest first. parent records,
// for each node reached, the node it was reached from.
func BFS(g Graph, start string) (order []string, parent map[string]string) {
	parent = map[string]string{}
	seen := map[string]bool{start: true}
	queue := []string{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)
		for _, next := range g[cur] {
			if !seen[next] {
				seen[next] = true // seen when queued, so it is queued once
				parent[next] = cur
				queue = append(queue, next)
			}
		}
	}
	return order, parent
}

// DFS visits the nodes reachable from start, going as deep as possible first.
func DFS(g Graph, start string) []string {
	var order []string
	seen := map[string]bool{}
	var visit func(n string)
	visit = func(n string) {
		if seen[n] {
			return
		}
		seen[n] = true
		order = append(order, n)
		for _, next := range g[n] {
			visit(next)
		}
	}
	visit(start)
	return order
}
