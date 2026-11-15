package main

// Graph maps each package to the packages it imports.
type Graph map[string][]string

// FindCycle returns one cycle as a path that starts and ends on the same node,
// such as [a b c a], or nil when the graph has no cycle.
func FindCycle(g Graph) []string {
	return nil
}
