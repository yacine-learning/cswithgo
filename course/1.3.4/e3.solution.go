package main

type Node struct {
	Key         int // expiry time of a certificate
	Left, Right *Node
}

// DeleteBelow removes every key < t and returns the new root.
func DeleteBelow(n *Node, t int) *Node {
	if n == nil {
		return nil
	}
	if n.Key < t {
		return DeleteBelow(n.Right, t) // n and its whole left side are expired
	}
	n.Left = DeleteBelow(n.Left, t)
	return n
}
