package main

import "fmt"

func main() {
	t := NewTrie()
	for _, w := range []string{"go", "goto", "gopher", "func", "for", "fallthrough"} {
		t.Insert(w)
	}
	for _, w := range []string{"go", "gop", "gopher", "fo", "for", "x"} {
		fmt.Printf("%-8s Contains: %-5v HasPrefix: %v\n", w, t.Contains(w), t.HasPrefix(w))
	}
	nodes := 0
	var count func(*Trie)
	count = func(n *Trie) {
		nodes++
		for _, c := range n.children {
			count(c)
		}
	}
	count(t)
	fmt.Println("\nnodes in the trie:", nodes, "(the letters of the 6 words, minus the shared prefixes, plus the root)")
}
