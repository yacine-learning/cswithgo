package main

type Node struct {
	Key         int
	Left, Right *Node
}

// Insert adds key to the tree rooted at n and returns the new root of that subtree.
func Insert(n *Node, key int) *Node {
	// hole: an empty subtree becomes a new node; otherwise insert on the correct side
	// and assign the result back to that side. Return n.
	return n
}

// Search reports whether key is in the tree, and how many nodes it looked at.
func Search(n *Node, key int) (found bool, visited int) {
	for n != nil {
		visited++
		if key == n.Key {
			return true, visited
		}
		if key < n.Key {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return false, visited
}

// Height is the number of nodes on the longest path from n down to a leaf.
func Height(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + max(Height(n.Left), Height(n.Right))
}
