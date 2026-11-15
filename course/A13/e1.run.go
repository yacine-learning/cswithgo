package main

import "fmt"

func both(g Graph, a, b string, cost int) {
	g[a] = append(g[a], Edge{b, cost})
	g[b] = append(g[b], Edge{a, cost})
}

func main() {
	// routers and the cost of each link (a slow link costs more). Link-state routing
	// protocols such as OSPF run exactly this on every router.
	g := Graph{}
	both(g, "A", "B", 10)
	both(g, "A", "C", 1)
	both(g, "C", "D", 2)
	both(g, "D", "B", 1)
	both(g, "B", "E", 4)
	both(g, "D", "E", 9)
	both(g, "E", "F", 1)
	for _, dst := range []string{"B", "E", "F", "A"} {
		path, cost, ok := ShortestPath(g, "A", dst)
		fmt.Printf("A → %s: %v, cost %d, reachable %v\n", dst, path, cost, ok)
	}
	fmt.Println("\nthe direct link A-B costs 10; A C D B costs 4. Fewest hops is not cheapest.")
}
