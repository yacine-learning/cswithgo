package main

import "fmt"

func main() {
	a := []int{7, 2, 9, 4, 1, 8, 5}
	p := partition(a, 0, len(a)-1)
	fmt.Printf("after one partition around 5: %v, pivot at index %d\n", a, p)
	b := []int{38, 27, 43, 3, 9, 82, 10, 3}
	QuickSort(b, 0, len(b)-1)
	fmt.Println("sorted:", b)
}
