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
	var certs *Node
	for _, day := range []int{120, 40, 300, 10, 90, 200, 365, 60} {
		certs = insert(certs, day)
	}
	fmt.Println("expiry days:", keys(certs))
	certs = DeleteBelow(certs, 100)
	fmt.Println("today is day 100, after cleanup:", keys(certs))
}
