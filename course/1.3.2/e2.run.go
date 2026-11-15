package main

import "fmt"

func inorder(n *Node, out *[]int) {
	if n == nil {
		return
	}
	inorder(n.Left, out)
	*out = append(*out, n.Key)
	inorder(n.Right, out)
}

func main() {
	var root *Node
	for _, k := range []int{4, 6, 7, 5, 2, 1, 3} {
		root = Insert(root, k)
	}
	var keys []int
	inorder(root, &keys)
	fmt.Println("keys in order:", keys)
	fmt.Println("height:", Height(root))
	for _, k := range []int{5, 9} {
		found, v := Search(root, k)
		fmt.Printf("search %d: found=%v after %d nodes\n", k, found, v)
	}
}
