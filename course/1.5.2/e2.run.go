package main

import "fmt"

func main() {
	a := []int{1, 1, 2, 3, 3, 3, 7, 9, 9}
	fmt.Println("before:", a)
	k := Dedup(a)
	fmt.Println("after: ", a[:k], " k =", k)
	fmt.Println("whole slice now:", a, "(what is after k does not matter)")
}
