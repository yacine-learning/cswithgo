package main

import "fmt"

func main() {
	var v Vec
	fmt.Println("push  len  cap  copies so far")
	for i := 1; i <= 20; i++ {
		before := v.Cap()
		v.Push(i)
		if v.Cap() != before {
			fmt.Printf("%-5d %-4d %-4d %d   <- grew\n", i, v.Len(), v.Cap(), v.Copies)
		}
	}
	var big Vec
	for i := 0; i < 1_000_000; i++ {
		big.Push(i)
	}
	fmt.Printf("\n1 000 000 pushes: %d copies, %.2f per element\n", big.Copies, float64(big.Copies)/1e6)

	fmt.Println("\nGo's own slices: capacity after each growth")
	var s []int
	last := -1
	for i := 0; i < 5000; i++ {
		s = append(s, i)
		if cap(s) != last {
			fmt.Print(cap(s), " ")
			last = cap(s)
		}
	}
	fmt.Println()
}
