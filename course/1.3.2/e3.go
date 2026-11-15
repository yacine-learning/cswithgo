package main

type Node struct {
	Key         int // when a geological period began, in millions of years (negative: ago)
	Left, Right *Node
}

// Floor returns the largest key <= t. found is false when every key is > t.
func Floor(n *Node, t int) (key int, found bool) {
	return 0, false
}
