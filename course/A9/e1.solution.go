package main

import "strconv"

// Node is a number (Op == 0) or an operator ('+', '-', '*', '/') with two operands.
type Node struct {
	Op          rune
	Num         int
	Left, Right *Node
}

// Eval computes the value of the expression: post-order.
func Eval(n *Node) int {
	if n.Op == 0 {
		return n.Num
	}
	l, r := Eval(n.Left), Eval(n.Right)
	switch n.Op {
	case '+':
		return l + r
	case '-':
		return l - r
	case '*':
		return l * r
	default:
		return l / r
	}
}

// Infix writes the expression with brackets around every operation: in-order.
// 3 * (4 + 1) is "(3 * (4 + 1))".
func Infix(n *Node) string {
	if n.Op == 0 {
		return strconv.Itoa(n.Num)
	}
	return "(" + Infix(n.Left) + " " + string(n.Op) + " " + Infix(n.Right) + ")"
}
