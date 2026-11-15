package main

import "fmt"

var visited int

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

func expect(r *Node, t int, want string) func() string {
	return func() string {
		r = DeleteBelow(r, t)
		if got := fmt.Sprint(keys(r)); got != want {
			return fmt.Sprintf("DeleteBelow(%d) left %s, expected %s", t, got, want)
		}
		return ""
	}
}

func main() {
	base := func() *Node { return build(120, 40, 300, 10, 90, 200, 365, 60) }
	try("remove the expired ones", expect(base(), 100, "[120 200 300 365]"))
	try("the root itself expires", expect(base(), 150, "[200 300 365]"))
	try("nothing expired", expect(base(), 5, "[10 40 60 90 120 200 300 365]"))
	try("everything expired", expect(base(), 1000, "[]"))
	try("t equal to a key keeps that key", expect(base(), 90, "[90 120 200 300 365]"))
	try("an empty tree", expect(nil, 10, "[]"))
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
