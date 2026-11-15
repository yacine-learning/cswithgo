package main

import (
	"fmt"
	"strings"
)

func show(h MinHeap) {
	fmt.Println("  slice:", []int(h))
	for level, start := 0, 0; start < len(h); level, start = level+1, 2*start+1 {
		end := min(2*start+1, len(h))
		fmt.Printf("  %s%v\n", strings.Repeat("  ", 3-level), []int(h[start:end]))
	}
}

func main() {
	var h MinHeap
	for _, p := range []int{5, 3, 8, 1, 9, 2, 7} {
		h.Push(p)
		fmt.Printf("push %d:\n", p)
		show(h)
	}
	fmt.Print("\npop all: ")
	for len(h) > 0 {
		fmt.Print(h.Pop(), " ")
	}
	fmt.Println()
}
