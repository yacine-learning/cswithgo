package main

import "fmt"

func main() {
	g := Build([]string{"web -> api", "api -> db", "api->cache", "worker -> db"})
	for _, n := range []string{"web", "api", "worker", "db", "cache"} {
		outs, ok := g[n]
		fmt.Printf("%-7s in graph: %-5v -> %v\n", n, ok, outs)
	}
}
