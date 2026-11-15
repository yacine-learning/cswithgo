package main

import "fmt"

func valid(n *Node, lo, hi int) bool {
	if n == nil {
		return true
	}
	if n.Key <= lo || n.Key >= hi {
		return false
	}
	return valid(n.Left, lo, n.Key) && valid(n.Right, n.Key, hi)
}

func count(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + count(n.Left) + count(n.Right)
}

func main() {
	try("insert into an empty tree", func() string {
		root := Insert(nil, 5)
		if root == nil || root.Key != 5 {
			return fmt.Sprintf("Insert(nil, 5) = %v, expected a node with key 5", root)
		}
		return ""
	})
	try("the shape of 4 6 7 5 2 1 3", func() string {
		var root *Node
		for _, k := range []int{4, 6, 7, 5, 2, 1, 3} {
			root = Insert(root, k)
		}
		if root == nil || root.Key != 4 || root.Left == nil || root.Left.Key != 2 || root.Right == nil || root.Right.Key != 6 {
			return "expected 4 at the root, 2 on its left, 6 on its right"
		}
		if root.Left.Left == nil || root.Left.Left.Key != 1 || root.Right.Right == nil || root.Right.Right.Key != 7 {
			return "expected 1 below 2 on the left, 7 below 6 on the right"
		}
		return ""
	})
	try("every node: smaller keys left, larger right", func() string {
		var root *Node
		for _, k := range []int{50, 30, 70, 20, 40, 60, 80, 35, 45, 65} {
			root = Insert(root, k)
		}
		if !valid(root, -1<<62, 1<<62) {
			return "a key is on the wrong side of one of its ancestors"
		}
		if count(root) != 10 {
			return fmt.Sprintf("the tree has %d nodes after 10 inserts", count(root))
		}
		return ""
	})
	try("inserting an existing key changes nothing", func() string {
		var root *Node
		for _, k := range []int{2, 1, 3, 2, 1} {
			root = Insert(root, k)
		}
		if count(root) != 3 {
			return fmt.Sprintf("inserting 2 1 3 2 1 gave %d nodes, expected 3", count(root))
		}
		return ""
	})
	try("search finds what was inserted", func() string {
		var root *Node
		for k := 0; k < 100; k += 3 {
			root = Insert(root, k)
		}
		for k := 0; k < 100; k++ {
			found, _ := Search(root, k)
			if found != (k%3 == 0) {
				return fmt.Sprintf("Search(%d) = %v", k, found)
			}
		}
		return ""
	})
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
