package main

import "fmt"

type Node struct {
	Val         string
	Left, Right *Node
}

func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Val, " ") // the node first
	PreOrder(n.Left)
	PreOrder(n.Right)
}

func InOrder(n *Node) {
	if n == nil {
		return
	}
	InOrder(n.Left)
	fmt.Print(n.Val, " ") // the node between its two sides
	InOrder(n.Right)
}

func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.Left)
	PostOrder(n.Right)
	fmt.Print(n.Val, " ") // the node last
}
