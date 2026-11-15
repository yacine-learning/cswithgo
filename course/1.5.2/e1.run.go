package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 6, 8, 9, 11, 14, 17}
	fmt.Println("a =", a, "target 23")
	i, j, ok := PairSum(a, 23)
	fmt.Printf("-> a[%d]+a[%d], found=%v\n", i, j, ok)

	nested := 0
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			nested++
		}
	}
	fmt.Printf("\nthe nested loop checks up to %d pairs; two pointers at most %d\n", nested, len(a)-1)
}
