package main

import "strings"

// Graph maps each node to the nodes it points to.
type Graph map[string][]string

// Build reads lines "A -> B" (spaces around the arrow are optional).
// Every node that appears becomes a key; a node with no outgoing edge has an empty list.
func Build(lines []string) Graph {
	g := Graph{}
	for _, line := range lines {
		parts := strings.Split(line, "->")
		a, b := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if _, ok := g[b]; !ok {
			g[b] = nil
		}
		g[a] = append(g[a], b)
	}
	return g
}
