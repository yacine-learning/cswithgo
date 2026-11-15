package main

import "fmt"

func main() {
	try("edges in input order", func() string {
		g := Build([]string{"fox -> rabbit", "fox -> mouse", "fox -> bird"})
		if fmt.Sprint(g["fox"]) != "[rabbit mouse bird]" {
			return fmt.Sprintf("g[\"fox\"] = %v, expected [rabbit mouse bird]", g["fox"])
		}
		return ""
	})
	try("a node only on the right is still a key", func() string {
		g := Build([]string{"rabbit -> grass"})
		outs, ok := g["grass"]
		if !ok {
			return "\"grass\" is not a key of the graph"
		}
		if len(outs) != 0 {
			return fmt.Sprintf("g[\"grass\"] = %v, expected no edges", outs)
		}
		return ""
	})
	try("spaces around the arrow are optional", func() string {
		g := Build([]string{"a->b", "  c   ->   d  "})
		if fmt.Sprint(g["a"]) != "[b]" || fmt.Sprint(g["c"]) != "[d]" {
			return fmt.Sprintf("g = %v, expected a -> [b], c -> [d]", g)
		}
		if _, ok := g["a "]; ok {
			return "a key has a trailing space: trim both sides"
		}
		return ""
	})
	try("the number of nodes", func() string {
		g := Build([]string{"fox -> rabbit", "rabbit -> grass", "owl -> mouse", "fox -> mouse"})
		if len(g) != 5 {
			return fmt.Sprintf("%d nodes, expected 5 (fox rabbit grass owl mouse)", len(g))
		}
		return ""
	})
}

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
