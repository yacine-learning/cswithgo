package main

import (
	"sort"
	"strings"
	"unicode"
)

// Count is how many times a word appears.
type Count struct {
	Word string
	N    int
}

// WordCounts counts the words of text (runs of letters, in lower case),
// most frequent first, then in alphabetical order.
func WordCounts(text string) []Count {
	counts := map[string]int{}
	for _, w := range strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) }) {
		counts[strings.ToLower(w)]++
	}
	out := make([]Count, 0, len(counts))
	for w, n := range counts {
		out = append(out, Count{w, n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].N != out[j].N {
			return out[i].N > out[j].N
		}
		return out[i].Word < out[j].Word
	})
	return out
}
