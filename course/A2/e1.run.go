package main

import "fmt"

func main() {
	fmt.Println(" n   value         calls (slow)   calls (memo)")
	for _, n := range []int{10, 20, 25, 30} {
		Calls = 0
		v := FibSlow(n)
		slow := Calls
		Calls = 0
		Fib(n, map[int]int{})
		fmt.Printf("%2d   %-12d  %-14d %d\n", n, v, slow, Calls)
	}
}
