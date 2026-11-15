package main

import "fmt"

func main() {
	try("edges in input order", func() string {
		g := Build([]string{"api -> db", "api -> cache", "api -> queue"})
		if fmt.Sprint(g["api"]) != "[db cache queue]" {
			return fmt.Sprintf("g[\"api\"] = %v, expected [db cache queue]", g["api"])
		}
		return ""
	})
	try("a node only on the right is still a key", func() string {
		g := Build([]string{"api -> db"})
		outs, ok := g["db"]
		if !ok {
			return "\"db\" is not a key of the graph"
		}
		if len(outs) != 0 {
			return fmt.Sprintf("g[\"db\"] = %v, expected no edges", outs)
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
		g := Build([]string{"web -> api", "api -> db", "worker -> db", "api -> cache"})
		if len(g) != 5 {
			return fmt.Sprintf("%d nodes, expected 5 (web api db worker cache)", len(g))
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
