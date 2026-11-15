package main

import "fmt"

func insert(n *Node, k int) *Node {
	if n == nil {
		return &Node{Key: k}
	}
	if k < n.Key {
		n.Left = insert(n.Left, k)
	} else {
		n.Right = insert(n.Right, k)
	}
	return n
}

func keys(n *Node) []int {
	if n == nil {
		return nil
	}
	return append(append(keys(n.Left), n.Key), keys(n.Right)...)
}

func main() {
	var root *Node
	for _, k := range []int{50, 30, 70, 20, 40, 60, 80} {
		root = insert(root, k)
	}
	fmt.Println("keys:", keys(root))
	for _, k := range []int{20, 30, 50, 99} {
		root = Delete(root, k)
		fmt.Printf("delete %d -> %v\n", k, keys(root))
	}
}
