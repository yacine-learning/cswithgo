package main

// Node is a number (Op == 0) or an operator ('+', '-', '*', '/') with two operands.
type Node struct {
	Op          rune
	Num         int
	Left, Right *Node
}

// Eval computes the value of the expression: post-order.
func Eval(n *Node) int {
	// hole
	return 0
}

// Infix writes the expression with brackets around every operation: in-order.
// 3 * (4 + 1) is "(3 * (4 + 1))".
func Infix(n *Node) string {
	// hole
	return ""
}
