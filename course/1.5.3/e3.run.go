package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("it was the best of times it was the worst of times")
	s, n := LongestUnique(words)
	fmt.Println(words)
	fmt.Printf("longest run without a repeated word: %v (start %d, length %d)\n", words[s:s+n], s, n)
}
