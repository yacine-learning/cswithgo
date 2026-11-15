package main

import "fmt"

func main() {
	var h MinHeap
	for _, p := range []int{5, 3, 8, 1, 9, 2, 7} {
		h.Push(p)
	}
	fmt.Println("after the pushes:", []int(h), " (h[0] should be 1)")
	fmt.Print("pop all: ")
	for len(h) > 0 {
		fmt.Print(h.Pop(), " ")
	}
	fmt.Println()
}
