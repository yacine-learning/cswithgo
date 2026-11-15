package main

import "fmt"

func main() {
	try("counts and order", func() string {
		got := WordCounts("the cat and the dog and the bird")
		want := []Count{{"the", 3}, {"and", 2}, {"bird", 1}, {"cat", 1}, {"dog", 1}}
		if fmt.Sprint(got) != fmt.Sprint(want) {
			return fmt.Sprintf("got %v, expected %v", got, want)
		}
		return ""
	})
	try("upper and lower case are the same word", func() string {
		got := WordCounts("Sea sea SEA")
		if fmt.Sprint(got) != "[{sea 3}]" {
			return fmt.Sprintf("got %v, expected [{sea 3}]", got)
		}
		return ""
	})
	try("punctuation and line breaks separate words", func() string {
		got := WordCounts("life, with its several powers;\nlife!")
		want := "[{life 2} {its 1} {powers 1} {several 1} {with 1}]"
		if fmt.Sprint(got) != want {
			return fmt.Sprintf("got %v, expected %s", got, want)
		}
		return ""
	})
	try("no words", func() string {
		if got := WordCounts(" ... ;; "); len(got) != 0 {
			return fmt.Sprintf("got %v for a text without letters", got)
		}
		return ""
	})
	try("the same result every run (no map order leaks out)", func() string {
		text := "a b c d e f g h a b c d e f g h i j k"
		first := fmt.Sprint(WordCounts(text))
		for i := 0; i < 20; i++ {
			if s := fmt.Sprint(WordCounts(text)); s != first {
				return "two calls on the same text gave different orders:\n  " + first + "\n  " + s
			}
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
