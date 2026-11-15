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

func build(keys ...int) *Node {
	var root *Node
	for _, k := range keys {
		root = insert(root, k)
	}
	return root
}

func expect(root *Node, t, want int, wantFound bool) func() string {
	return func() string {
		k, ok := Floor(root, t)
		if ok != wantFound || (ok && k != want) {
			if wantFound {
				return fmt.Sprintf("Floor(%d) = %d, %v; expected %d, true", t, k, ok, want)
			}
			return fmt.Sprintf("Floor(%d) = %d, %v; expected found=false", t, k, ok)
		}
		return ""
	}
}

func main() {
	root := build(1400, 900, 1700, 1100, 1530, 2000)
	try("exactly a key", expect(root, 1100, 1100, true))
	try("between two keys", expect(root, 1432, 1400, true))
	try("the answer is deep on the left side", expect(root, 1200, 1100, true))
	try("the answer is in a right subtree", expect(root, 1600, 1530, true))
	try("after the last key", expect(root, 2359, 2000, true))
	try("before the first key", expect(root, 800, 0, false))
	try("an empty tree", expect(nil, 5, 0, false))
	try("a straight-line tree", expect(build(1, 2, 3, 4, 5, 6, 7, 8, 9), 6, 6, true))
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
