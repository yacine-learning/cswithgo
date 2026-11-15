package main

// Calls counts the calls made, to compare the two versions.
var Calls int

// FibSlow is the textbook definition. It computes the same values again and again.
func FibSlow(n int) int {
	Calls++
	if n < 2 {
		return n
	}
	return FibSlow(n-1) + FibSlow(n-2)
}

// Fib computes each value once and remembers it in memo.
func Fib(n int, memo map[int]int) int {
	Calls++
	if n < 2 {
		return n
	}
	// hole 1: if memo already has n, return it

	v := Fib(n-1, memo) + Fib(n-2, memo)
	// hole 2: remember v for n

	return v
}
