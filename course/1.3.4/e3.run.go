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
	var obs *Node
	for _, day := range []int{120, 40, 300, 10, 90, 200, 365, 60} {
		obs = insert(obs, day)
	}
	fmt.Println("observation days:", keys(obs))
	obs = DeleteBelow(obs, 100)
	fmt.Println("keep day 100 onwards:", keys(obs))
}
