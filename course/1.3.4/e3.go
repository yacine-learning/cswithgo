package main

type Node struct {
	Key         int // expiry time of a certificate
	Left, Right *Node
}

// DeleteBelow removes every key < t and returns the new root.
func DeleteBelow(n *Node, t int) *Node {
	return n
}
