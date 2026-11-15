package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := `In the beginning God created the heaven and the earth. And the earth was without form,
and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face
of the waters. And God said, Let there be light: and there was light.`
	counts := map[string]int{}
	for _, w := range strings.FieldsFunc(strings.ToLower(text), func(r rune) bool { return !unicode.IsLetter(r) }) {
		counts[w]++
	}
	fmt.Println("top 5:", TopK(counts, 5))
}
