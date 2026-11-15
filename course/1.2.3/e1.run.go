package main

import "fmt"

func main() {
	var s Stack
	pairs := map[string]string{")": "(", "]": "[", "}": "{"}
	for _, r := range "{[()]}" {
		c := string(r)
		if open, isCloser := pairs[c]; isCloser {
			top, _ := s.Pop()
			fmt.Printf("read %s  pop %s  (matches %s: %v)  stack %v\n", c, top, open, top == open, s.items)
		} else {
			s.Push(c)
			fmt.Printf("read %s  push       stack %v\n", c, s.items)
		}
	}
	_, ok := s.Pop()
	fmt.Println("\nPop on the empty stack: ok =", ok)
}
