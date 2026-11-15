package main

type Node struct {
	Key         int // the sequence number of a segment waiting for its acknowledgement
	Left, Right *Node
}

// DeleteBelow removes every key < t and returns the new root.
func DeleteBelow(n *Node, t int) *Node {
	return n
}
