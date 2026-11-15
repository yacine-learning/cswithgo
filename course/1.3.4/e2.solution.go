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
	// found it
	if n.Left == nil {
		return n.Right // no child (nil), or only a right child: it takes n's place
	}
	if n.Right == nil {
		return n.Left // only a left child
	}
	// two children: copy in the successor, then delete the successor from the right side
	n.Key = Min(n.Right)
	n.Right = Delete(n.Right, n.Key)
	return n
}
