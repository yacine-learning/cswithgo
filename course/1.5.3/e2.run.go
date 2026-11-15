package main

import "fmt"

func main() {
	l := &Limiter{Limit: 3, Window: 10}
	for _, t := range []int{0, 1, 2, 3, 9, 10, 11, 12, 13, 25} {
		fmt.Printf("t=%2d  allowed=%v\n", t, l.Allow(t))
	}
}
