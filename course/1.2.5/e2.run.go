package main

import "fmt"

func main() {
	l := FromValues(1, 2, 3, 4, 5)
	fmt.Println("before: ", l.Values())
	l.Reverse()
	fmt.Println("after:  ", l.Values())
}
