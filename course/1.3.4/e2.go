package main

type Node struct {
	Key         int
	Left, Right *Node
}

// Min returns the smallest key of a non-empty tree: keep going left.
func Min(n *Node) int {
	for n.Left != nil {
		n = n.Left
	}
	return n.Key
}

// Max returns the largest key of a non-empty tree: keep going right.
func Max(n *Node) int {
	for n.Right != nil {
		n = n.Right
	}
	return n.Key
}

// Delete removes key from the tree rooted at n and returns the new root.
func Delete(n *Node, key int) *Node {
	if n == nil {
		return nil // not found: nothing changes
	}
	if key < n.Key {
		n.Left = Delete(n.Left, key) // the left subtree may get a new root: reconnect it
		return n
	}
	if key > n.Key {
		n.Right = Delete(n.Right, key)
		return n
	}
	// found it: n is the node to delete
	// hole 1: no left child -> its right child (maybe nil) takes its place

	// hole 2: no right child -> its left child takes its place

	// hole 3: two children -> copy in the minimum of the right subtree, then delete that key there

	return n
}
