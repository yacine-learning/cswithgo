package main

import "fmt"

// isCycle checks that p is a real cycle of g: closed, every step an edge.
func isCycle(g Graph, p []string) string {
	if len(p) < 2 || p[0] != p[len(p)-1] {
		return fmt.Sprintf("%v does not start and end on the same node", p)
	}
	for i := 0; i+1 < len(p); i++ {
		found := false
		for _, n := range g[p[i]] {
			found = found || n == p[i+1]
		}
		if !found {
			return fmt.Sprintf("%v uses %s -> %s, which is not an edge", p, p[i], p[i+1])
		}
	}
	return ""
}

func main() {
	try("no cycle in a DAG", func() string {
		g := Graph{"puppy": {"young", "dog"}, "kitten": {"young", "cat"}, "dog": {"animal"}, "cat": {"animal"}, "young": nil, "animal": nil}
		if c := FindCycle(g); c != nil {
			return fmt.Sprintf("FindCycle = %v on a graph without a cycle (two words sharing a word in their definitions is not a cycle)", c)
		}
		return ""
	})
	try("a three-node cycle", func() string {
		g := Graph{"huge": {"big"}, "big": {"large", "size"}, "large": {"great"}, "great": {"big"}, "size": nil}
		c := FindCycle(g)
		if c == nil {
			return "no cycle found in big -> large -> great -> big"
		}
		return isCycle(g, c)
	})
	try("a word defined by itself", func() string {
		g := Graph{"a": {"a"}}
		c := FindCycle(g)
		if c == nil {
			return "no cycle found in a -> a"
		}
		return isCycle(g, c)
	})
	try("the cycle is only in a second part of the graph", func() string {
		g := Graph{"a": {"b"}, "b": nil, "x": {"y"}, "y": {"z"}, "z": {"x"}}
		c := FindCycle(g)
		if c == nil {
			return "no cycle found in x -> y -> z -> x (the DFS must start from every node)"
		}
		return isCycle(g, c)
	})
	try("a diamond is not a cycle", func() string {
		g := Graph{"a": {"b", "c"}, "b": {"d"}, "c": {"d"}, "d": nil}
		if c := FindCycle(g); c != nil {
			return fmt.Sprintf("FindCycle = %v on a -> b -> d, a -> c -> d: d is reached twice, but that is not a cycle", c)
		}
		return ""
	})
	try("an empty graph", func() string {
		if c := FindCycle(Graph{}); c != nil {
			return fmt.Sprintf("FindCycle = %v on an empty graph", c)
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
