package main

import "fmt"

func main() {
	a := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println("a =", a)
	for _, t := range []int{23, 2, 91, 50} {
		fmt.Printf("\nsearch %d:\n", t)
		fmt.Println("  ->", Search(a, t))
	}
}
