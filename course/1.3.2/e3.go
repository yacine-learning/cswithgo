package main

type Node struct {
	Key         int // the time a config version took effect
	Left, Right *Node
}

// Floor returns the largest key <= t. found is false when every key is > t.
func Floor(n *Node, t int) (key int, found bool) {
	return 0, false
}
