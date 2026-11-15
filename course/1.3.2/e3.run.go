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

func main() {
	var versions *Node
	for _, t := range []int{1400, 900, 1700, 1100, 1530, 2000} { // HHMM when each config went live
		versions = insert(versions, t)
	}
	for _, t := range []int{800, 900, 1432, 1600, 2359} {
		k, ok := Floor(versions, t)
		fmt.Printf("at %04d the config active since %04d (found=%v)\n", t, k, ok)
	}
}
