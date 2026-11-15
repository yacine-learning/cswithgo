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

func build(ks ...int) *Node {
	var r *Node
	for _, k := range ks {
		r = insert(r, k)
	}
	return r
}

func keys(n *Node) []int {
	if n == nil {
		return nil
	}
	return append(append(keys(n.Left), n.Key), keys(n.Right)...)
}

func valid(n *Node, lo, hi int) bool {
	return n == nil || (n.Key > lo && n.Key < hi && valid(n.Left, lo, n.Key) && valid(n.Right, n.Key, hi))
}

func expect(r *Node, k int, want string) func() string {
	return func() string {
		r = Delete(r, k)
		if got := fmt.Sprint(keys(r)); got != want {
			return fmt.Sprintf("after Delete(%d): keys %s, expected %s", k, got, want)
		}
		if !valid(r, -1<<62, 1<<62) {
			return fmt.Sprintf("after Delete(%d) a key is on the wrong side", k)
		}
		return ""
	}
}

func main() {
	try("a leaf", expect(build(50, 30, 70, 20), 20, "[30 50 70]"))
	try("a node with only a right child", expect(build(50, 30, 35), 30, "[35 50]"))
	try("a node with only a left child", expect(build(50, 30, 20), 30, "[20 50]"))
	try("a node with two children", expect(build(50, 30, 70, 20, 40, 60, 80, 65), 70, "[20 30 40 50 60 65 80]"))
	try("the root with two children", expect(build(50, 30, 70, 20, 40, 60, 80), 50, "[20 30 40 60 70 80]"))
	try("the root alone", func() string {
		if r := Delete(build(5), 5); r != nil {
			return fmt.Sprintf("deleting the only key left a tree with %v", keys(r))
		}
		return ""
	})
	try("a key that is not there", expect(build(2, 1, 3), 9, "[1 2 3]"))
	try("delete everything, one by one", func() string {
		r := build(50, 30, 70, 20, 40, 60, 80, 35, 45, 65)
		for _, k := range []int{50, 20, 65, 30, 80, 45, 70, 35, 60, 40} {
			r = Delete(r, k)
			if !valid(r, -1<<62, 1<<62) {
				return fmt.Sprintf("after Delete(%d) a key is on the wrong side", k)
			}
		}
		if r != nil {
			return fmt.Sprintf("after deleting all ten keys, %v remain", keys(r))
		}
		return ""
	})
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
