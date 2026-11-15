package main

import "fmt"

func main() {
	try("edges in input order", func() string {
		g := Build([]string{"main -> fmt", "main -> lexer", "main -> sort"})
		if fmt.Sprint(g["main"]) != "[fmt lexer sort]" {
			return fmt.Sprintf("g[\"main\"] = %v, expected [fmt lexer sort]", g["main"])
		}
		return ""
	})
	try("a node only on the right is still a key", func() string {
		g := Build([]string{"fmt -> io"})
		outs, ok := g["io"]
		if !ok {
			return "\"io\" is not a key of the graph"
		}
		if len(outs) != 0 {
			return fmt.Sprintf("g[\"io\"] = %v, expected no edges", outs)
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
		g := Build([]string{"main -> fmt", "fmt -> io", "parser -> lexer", "main -> lexer"})
		if len(g) != 5 {
			return fmt.Sprintf("%d nodes, expected 5 (main fmt io parser lexer)", len(g))
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
