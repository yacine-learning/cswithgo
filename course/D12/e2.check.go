package main

import (
	"fmt"
	"strings"
)

func main() {
	t := NewTrie()
	for _, w := range []string{"copy", "const", "continue", "close", "complex", "case", "go", "goto", "cap"} {
		t.Insert(w)
	}
	show := func(s []string) string { return "[" + strings.Join(s, " ") + "]" }
	try("words with a prefix, in order", func() string {
		for i := 0; i < 20; i++ { // map order changes from run to run
			got := show(Complete(t, "co"))
			if got != "[complex const continue copy]" {
				return fmt.Sprintf(`Complete(t, "co") = %s, expected [complex const continue copy]`, got)
			}
		}
		return ""
	})
	try("the prefix itself is a word", func() string {
		if got := show(Complete(t, "go")); got != "[go goto]" {
			return fmt.Sprintf(`Complete(t, "go") = %s, expected [go goto]`, got)
		}
		return ""
	})
	try("a whole word with nothing after it", func() string {
		if got := show(Complete(t, "copy")); got != "[copy]" {
			return fmt.Sprintf(`Complete(t, "copy") = %s, expected [copy]`, got)
		}
		return ""
	})
	try("no word starts with the prefix", func() string {
		if got := Complete(t, "x"); len(got) != 0 {
			return fmt.Sprintf(`Complete(t, "x") = %s, expected nothing`, show(got))
		}
		return ""
	})
	try("the empty prefix gives every word", func() string {
		got := show(Complete(t, ""))
		want := "[cap case close complex const continue copy go goto]"
		if got != want {
			return fmt.Sprintf(`Complete(t, "") = %s, expected %s`, got, want)
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
