package main

import "fmt"

func main() {
	g := Build([]string{"fox -> rabbit", "rabbit -> grass", "owl->mouse", "fox -> mouse"})
	for _, n := range []string{"fox", "rabbit", "owl", "mouse", "grass"} {
		outs, ok := g[n]
		fmt.Printf("%-7s in graph: %-5v -> %v\n", n, ok, outs)
	}
}
