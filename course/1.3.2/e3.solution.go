package main

type Node struct {
	Key         int // the time a config version took effect
	Left, Right *Node
}

// Floor returns the largest key <= t. found is false when every key is > t.
func Floor(n *Node, t int) (key int, found bool) {
	for n != nil {
		if n.Key <= t {
			key, found = n.Key, true // a candidate; a larger one can only be on the right
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return key, found
}
