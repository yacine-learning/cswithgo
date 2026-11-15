package main

import "fmt"

func main() {
	// routers; links go both ways
	g := Graph{}
	link := func(a, b string) { g[a] = append(g[a], b); g[b] = append(g[b], a) }
	link("A", "B")
	link("A", "C")
	link("B", "D")
	link("C", "D")
	link("D", "E")
	link("E", "F")
	link("C", "F")
	fmt.Println("  A - B")
	fmt.Println("  |   |")
	fmt.Println("  C - D - E")
	fmt.Println("   \\      |")
	fmt.Println("     ----- F")
	order, parent := BFS(g, "A")
	fmt.Println("\nBFS from A:", order, " (by distance)")
	fmt.Println("DFS from A:", DFS(g, "A"), " (deep first)")
	path := []string{"E"}
	for n := "E"; n != "A"; n = parent[n] {
		path = append([]string{parent[n]}, path...)
	}
	fmt.Println("\nfewest hops A to E:", path)
}
