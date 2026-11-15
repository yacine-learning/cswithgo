package main

import "fmt"

func main() {
	r := NewRing(3)
	fmt.Println("enqueue a, b, c:", r.Enqueue("a"), r.Enqueue("b"), r.Enqueue("c"))
	fmt.Println("enqueue d (full):", r.Enqueue("d"))
	x, ok := r.Dequeue()
	fmt.Println("dequeue:", x, ok)
	fmt.Println("enqueue d:", r.Enqueue("d"))
	for r.Len() > 0 {
		x, _ = r.Dequeue()
		fmt.Println("dequeue:", x)
	}
	_, ok = r.Dequeue()
	fmt.Println("dequeue on empty: ok =", ok)
}
