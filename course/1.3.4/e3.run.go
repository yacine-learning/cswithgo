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
	var sent *Node
	for _, seq := range []int{120, 40, 300, 10, 90, 200, 365, 60} {
		sent = insert(sent, seq)
	}
	fmt.Println("segments not yet acknowledged:", keys(sent))
	sent = DeleteBelow(sent, 100) // the receiver says: I have everything before 100
	fmt.Println("after ACK 100:", keys(sent))
}
