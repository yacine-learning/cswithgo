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
	// a few Go packages and what they import (simplified)
	g.AddEdge("main", "fmt")
	g.AddEdge("main", "strings")
	g.AddEdge("fmt", "strconv")
	g.AddEdge("fmt", "io")
	g.AddEdge("strings", "unicode")
	g.AddEdge("strconv", "unicode")
	fmt.Println("imports:")
	show(g)
	fmt.Println("\nimported by (the reverse graph):")
	show(g.Reverse())
}
