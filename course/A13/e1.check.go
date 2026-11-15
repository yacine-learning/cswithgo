package main

import (
	"fmt"
	"strings"
)

func link(g Graph, a, b string, c int) { g[a] = append(g[a], Edge{b, c}) }

func main() {
	show := func(p []string) string { return strings.Join(p, " ") }
	try("cheaper through a detour", func() string {
		g := Graph{}
		link(g, "A", "B", 10)
		link(g, "A", "C", 1)
		link(g, "C", "B", 2)
		p, c, ok := ShortestPath(g, "A", "B")
		if !ok || c != 3 || show(p) != "A C B" {
			return fmt.Sprintf("A→B 10, A→C 1, C→B 2: got path [%s] cost %d ok %v, expected [A C B] cost 3", show(p), c, ok)
		}
		return ""
	})
	try("a better route found later", func() string {
		g := Graph{}
		link(g, "A", "B", 5)
		link(g, "A", "C", 1)
		link(g, "C", "D", 1)
		link(g, "D", "B", 1)
		link(g, "B", "E", 1)
		p, c, ok := ShortestPath(g, "A", "E")
		if !ok || c != 4 || show(p) != "A C D B E" {
			return fmt.Sprintf("B is first reached at cost 5, then at 3 through C and D: A→E got [%s] cost %d ok %v, expected [A C D B E] cost 4", show(p), c, ok)
		}
		return ""
	})
	try("from a node to itself", func() string {
		g := Graph{}
		link(g, "A", "B", 1)
		p, c, ok := ShortestPath(g, "A", "A")
		if !ok || c != 0 || show(p) != "A" {
			return fmt.Sprintf("A→A: got [%s] cost %d ok %v, expected [A] cost 0 true", show(p), c, ok)
		}
		return ""
	})
	try("unreachable", func() string {
		g := Graph{}
		link(g, "A", "B", 1)
		link(g, "C", "D", 1)
		if _, _, ok := ShortestPath(g, "A", "D"); ok {
			return "A→D with no route between them: ok = true"
		}
		if _, _, ok := ShortestPath(g, "B", "A"); ok {
			return "links are one-way: B→A has no route, but ok = true"
		}
		return ""
	})
	try("zero-cost links", func() string {
		g := Graph{}
		link(g, "A", "B", 0)
		link(g, "B", "C", 0)
		link(g, "A", "C", 1)
		if _, c, ok := ShortestPath(g, "A", "C"); !ok || c != 0 {
			return fmt.Sprintf("A→B 0, B→C 0, A→C 1: A→C cost %d, expected 0", c)
		}
		return ""
	})
	try("a 200 × 200 grid", func() string {
		n := 200
		g := Graph{}
		name := func(r, c int) string { return fmt.Sprintf("%d,%d", r, c) }
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if r+1 < n {
					link(g, name(r, c), name(r+1, c), 1)
					link(g, name(r+1, c), name(r, c), 1)
				}
				if c+1 < n {
					link(g, name(r, c), name(r, c+1), 1)
					link(g, name(r, c+1), name(r, c), 1)
				}
			}
		}
		p, c, ok := ShortestPath(g, name(0, 0), name(n-1, n-1))
		if !ok || c != 2*(n-1) || len(p) != 2*n-1 {
			return fmt.Sprintf("corner to corner: cost %d with %d nodes, expected %d with %d", c, len(p), 2*(n-1), 2*n-1)
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
