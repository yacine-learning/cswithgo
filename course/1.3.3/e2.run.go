package main

import "fmt"

func num(v int) *Node             { return &Node{Num: v} }
func op(o rune, l, r *Node) *Node { return &Node{Op: o, Left: l, Right: r} }

func main() {
	e := op('*', num(3), op('+', num(4), op('-', num(5), num(1))))
	fmt.Println(Infix(e), "=", Eval(e))
}
