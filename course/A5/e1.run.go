package main

import "fmt"

func main() {
	c := &Counter{Window: 10}
	for _, t := range []int{0, 1, 2, 3, 9, 10, 11, 12, 13, 25} {
		fmt.Printf("click at t=%2d: %d clicks in the last 10 s\n", t, c.Click(t))
	}
}
