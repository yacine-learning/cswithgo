package main

import (
	"fmt"
	"sort"
)

func show(g Graph) {
	names := make([]string, 0, len(g))
	for n := range g {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		outs := append([]string(nil), g[n]...)
		sort.Strings(outs)
		fmt.Printf("  %-8s -> %v\n", n, outs)
	}
}

func main() {
	g := Graph{}
	g.AddEdge("web", "api")
	g.AddEdge("api", "db")
	g.AddEdge("api", "cache")
	g.AddEdge("worker", "db")
	g.AddEdge("worker", "queue")
	g.AddEdge("api", "queue")
	fmt.Println("depends on:")
	show(g)
	fmt.Println("\nneeded by (the reverse graph):")
	show(g.Reverse())
}
