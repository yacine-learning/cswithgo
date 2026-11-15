package main

import (
	"fmt"
	"strings"
)

func main() {
	cases := []struct {
		a, b string
		want int
	}{
		{"kitten", "sitting", 3}, {"flaw", "lawn", 2}, {"intention", "execution", 5},
		{"go", "go", 0}, {"", "abc", 3}, {"abc", "", 3}, {"", "", 0}, {"a", "b", 1}, {"café", "cafe", 1},
	}
	try("known distances", func() string {
		for _, c := range cases {
			if got := EditDistance(c.a, c.b); got != c.want {
				return fmt.Sprintf("EditDistance(%q, %q) = %d, expected %d", c.a, c.b, got, c.want)
			}
		}
		return ""
	})
	try("symmetric", func() string {
		for _, c := range cases {
			if EditDistance(c.a, c.b) != EditDistance(c.b, c.a) {
				return fmt.Sprintf("EditDistance(%q, %q) differs from EditDistance(%q, %q)", c.a, c.b, c.b, c.a)
			}
		}
		return ""
	})
	// last: without a table this cannot finish
	try("two texts of 2 000 letters", func() string {
		a, b := strings.Repeat("ab", 1000), strings.Repeat("ba", 1000)
		if got := EditDistance(a, b); got != 2 {
			return fmt.Sprintf("abab…ab against baba…ba (2 000 letters): %d, expected 2 (delete the first a, add one at the end)", got)
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
