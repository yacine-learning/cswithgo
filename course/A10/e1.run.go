package main

import "fmt"

func main() {
	g := Graph{}
	link := func(a, b string) { g[a] = append(g[a], b); g[b] = append(g[b], a) }
	link("A", "B")
	link("A", "C")
	link("B", "D")
	link("C", "D")
	link("D", "E")
	link("E", "F")
	link("C", "F")
	g["Z"] = nil
	for _, to := range []string{"A", "D", "E", "F", "Z"} {
		fmt.Printf("A to %s: %v\n", to, ShortestPath(g, "A", to))
	}
}
