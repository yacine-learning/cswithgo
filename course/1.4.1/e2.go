package main

// Graph maps each node to the nodes it points to.
type Graph map[string][]string

// Build reads lines "A -> B" (spaces around the arrow are optional).
// Every node that appears becomes a key; a node with no outgoing edge has an empty list.
func Build(lines []string) Graph {
	g := Graph{}
	for _, line := range lines {
		// hole: split the line into a and b, make sure b is a key, then add the edge a -> b
		_ = line // delete this line once you use line
	}
	return g
}
