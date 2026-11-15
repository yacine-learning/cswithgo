package main

import "fmt"

func network() Graph {
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
	return g
}

func expect(g Graph, from, to, want string) func() string {
	return func() string {
		if got := fmt.Sprint(ShortestPath(g, from, to)); got != want {
			return fmt.Sprintf("ShortestPath(%s, %s) = %s, expected %s", from, to, got, want)
		}
		return ""
	}
}

func main() {
	g := network()
	try("one hop", expect(g, "A", "B", "[A B]"))
	try("two hops", expect(g, "A", "D", "[A B D]"))
	try("the shortcut through C", expect(g, "A", "F", "[A C F]"))
	try("three hops", expect(g, "A", "E", "[A B D E]"))
	try("from a node to itself", expect(g, "A", "A", "[A]"))
	try("unreachable", expect(g, "A", "Z", "[]"))
	try("a cycle does not loop forever", expect(Graph{"a": {"b"}, "b": {"c"}, "c": {"a", "d"}}, "a", "d", "[a b c d]"))
	try("a long chain", func() string {
		g := Graph{}
		for i := 0; i < 1000; i++ {
			g[fmt.Sprint(i)] = []string{fmt.Sprint(i + 1)}
		}
		p := ShortestPath(g, "0", "1000")
		if len(p) != 1001 || p[0] != "0" || p[1000] != "1000" {
			return fmt.Sprintf("path of %d nodes from 0 to 1000, expected 1001", len(p))
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
