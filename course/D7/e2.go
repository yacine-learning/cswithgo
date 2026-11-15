package main

type Node struct {
	Key         int // the address where a function's machine code begins
	Left, Right *Node
}

// Floor returns the largest key <= t. found is false when every key is > t.
func Floor(n *Node, t int) (key int, found bool) {
	return 0, false
}
