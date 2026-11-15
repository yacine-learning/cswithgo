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
		fmt.Printf("  %-12s -> %v\n", n, outs)
	}
}

func main() {
	g := Graph{}
	g.AddEdge("fox", "rabbit")
	g.AddEdge("fox", "mouse")
	g.AddEdge("owl", "mouse")
	g.AddEdge("rabbit", "grass")
	g.AddEdge("mouse", "seeds")
	g.AddEdge("mouse", "grass")
	fmt.Println("eats:")
	show(g)
	fmt.Println("\neaten by (the reverse graph):")
	show(g.Reverse())
}
