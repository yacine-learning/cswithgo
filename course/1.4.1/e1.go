package main

// Graph maps each node to the nodes it points to (its outgoing edges).
type Graph map[string][]string

// AddEdge adds a -> b. Both nodes become keys, so a node with no outgoing
// edge still appears in the graph.
func (g Graph) AddEdge(a, b string) {
	if _, ok := g[b]; !ok {
		g[b] = nil
	}
	g[a] = append(g[a], b)
}

// Reverse returns the graph with every edge turned around: "who points to me".
func (g Graph) Reverse() Graph {
	r := Graph{}
	for a, outs := range g {
		if _, ok := r[a]; !ok {
			r[a] = nil
		}
		for _, b := range outs {
			r.AddEdge(b, a)
		}
	}
	return r
}
