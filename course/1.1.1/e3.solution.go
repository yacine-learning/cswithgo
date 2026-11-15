package main

// Each function returns how many steps it takes for a given n.
// Read them, then fill Answers with Constant, Logarithmic, Linear or Quadratic.

func stepsA(n int) int {
	steps := 0
	for i := n; i > 1; i /= 2 {
		steps++
	}
	return steps
}

func stepsB(n int) int {
	steps := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			steps++
		}
	}
	return steps
}

func stepsC(n int) int {
	steps := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			steps++
		}
	}
	return steps
}

func stepsD(n int) int {
	steps := 0
	for i := 0; i < n && i < 100; i++ {
		steps++
	}
	return steps
}

type Growth string

const (
	Unknown     Growth = "?"
	Constant    Growth = "O(1)"
	Logarithmic Growth = "O(log n)"
	Linear      Growth = "O(n)"
	Quadratic   Growth = "O(n²)"
)

var Answers = map[string]Growth{
	"A": Logarithmic,
	"B": Linear,
	"C": Quadratic,
	"D": Constant,
}
