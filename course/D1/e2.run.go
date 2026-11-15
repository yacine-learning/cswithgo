package main

import "fmt"

func main() {
	// a path through a maze, as a list of room numbers, built with append: it has spare room
	prefix := make([]int, 0, 10)
	prefix = append(prefix, 1, 4, 7)
	left := Extend(prefix, 8)  // explore the left door...
	right := Extend(prefix, 9) // ...then the right door
	fmt.Println("left: ", left)
	fmt.Println("right:", right)
	fmt.Println("prefix:", prefix)
}
