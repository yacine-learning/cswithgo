package main

import (
	"fmt"
	"strings"
)

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

// show prints the tree sideways: the root on the left, right subtree above.
func show(n *Node, depth int) {
	if n == nil {
		return
	}
	show(n.Right, depth+1)
	fmt.Printf("%s%d\n", strings.Repeat("    ", depth), n.Key)
	show(n.Left, depth+1)
}

func main() {
	var root *Node
	for _, k := range []int{50, 30, 70, 20, 40, 60, 80, 65} {
		root = insert(root, k)
	}
	fmt.Println("the tree (root on the left, read it tilted):")
	show(root, 0)
	fmt.Println("min", Min(root), "max", Max(root))
	for _, step := range []struct {
		k    int
		what string
	}{{20, "a leaf"}, {60, "one child (65)"}, {30, "now a leaf too"}, {50, "the root, two children: 60s successor 65 takes its place"}} {
		root = Delete(root, step.k)
		fmt.Printf("\ndelete %d (%s):\n", step.k, step.what)
		show(root, 0)
	}
}
