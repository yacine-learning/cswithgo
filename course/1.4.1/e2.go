package main

// Graph maps each node to the nodes it points to.
type Graph map[string][]string

// Build reads lines "A -> B" (spaces around the arrow are optional).
// Every node that appears becomes a key; a node with no outgoing edge has an empty list.
func Build(lines []string) Graph {
	g := Graph{}
	for _, line := range lines {
		// hole: split the line into a and b, make sure b is a key, then add the edge a -> b
		// Go refuses to compile a variable that is never used. This line "uses" line
		// by throwing the value away, so the file compiles before you fill the holes.
		// Delete it once your code uses line.
		_ = line
	}
	return g
}
