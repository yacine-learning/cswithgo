package main

import "fmt"

func main() {
	words := []string{"go", "goto", "gopher", "func", "for", "café"}
	t := NewTrie()
	for _, w := range words {
		t.Insert(w)
	}
	try("every inserted word is found", func() string {
		for _, w := range words {
			if !t.Contains(w) {
				return fmt.Sprintf("after inserting %q, Contains(%q) returned false", w, w)
			}
		}
		return ""
	})
	try("a prefix is not a word", func() string {
		for _, w := range []string{"g", "gop", "fun", "caf"} {
			if t.Contains(w) {
				return fmt.Sprintf("Contains(%q) returned true, but only longer words start with it", w)
			}
		}
		return ""
	})
	try("HasPrefix", func() string {
		for _, p := range []string{"g", "gop", "gopher", "fu", "caf"} {
			if !t.HasPrefix(p) {
				return fmt.Sprintf("HasPrefix(%q) returned false", p)
			}
		}
		for _, p := range []string{"x", "gox", "gophers", "fo r"} {
			if t.HasPrefix(p) {
				return fmt.Sprintf("HasPrefix(%q) returned true: no word starts with it", p)
			}
		}
		return ""
	})
	try("words that share a prefix keep each other", func() string {
		u := NewTrie()
		u.Insert("go")
		u.Insert("gopher")
		u.Insert("goto")
		for _, w := range []string{"go", "gopher", "goto"} {
			if !u.Contains(w) {
				return fmt.Sprintf("inserted go, gopher, goto: Contains(%q) returned false", w)
			}
		}
		return ""
	})
	try("empty trie", func() string {
		if NewTrie().Contains("go") {
			return "an empty trie Contains(\"go\")"
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
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
