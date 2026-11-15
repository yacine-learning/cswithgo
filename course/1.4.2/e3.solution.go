package main

import "sort"

// Graph maps each word to the words used in its definition.
type Graph map[string][]string

const (
	notVisited = iota
	inProgress // on the current path
	done
)

// FindCycle returns one cycle as a path that starts and ends on the same node,
// such as [a b c a], or nil when the graph has no cycle.
func FindCycle(g Graph) []string {
	state := map[string]int{}
	var path, cycle []string
	var visit func(n string) bool
	visit = func(n string) bool {
		state[n] = inProgress
		path = append(path, n)
		for _, next := range g[n] {
			switch state[next] {
			case inProgress: // an edge back to the current path: a cycle
				for i, p := range path {
					if p == next {
						cycle = append(append([]string(nil), path[i:]...), next)
						return true
					}
				}
			case notVisited:
				if visit(next) {
					return true
				}
			}
		}
		path = path[:len(path)-1]
		state[n] = done
		return false
	}
	names := make([]string, 0, len(g))
	for n := range g {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		if state[n] == notVisited && visit(n) {
			return cycle
		}
	}
	return nil
}
