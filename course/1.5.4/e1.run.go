package main

import "fmt"

func main() {
	ts := []int{0, 5, 30, 31, 31, 32, 33, 60, 61, 300, 301, 302, 303, 304, 305, 306, 900}
	fmt.Println("meteors at:", ts)
	fmt.Println("between 30 and 60:", Count(ts, 30, 60))
	s, n := Busiest(ts, 10)
	fmt.Printf("busiest 10 s: from %d, %d meteors\n", s, n)
	fmt.Println("silences of 120 s or more:", Silences(ts, 120))
}
