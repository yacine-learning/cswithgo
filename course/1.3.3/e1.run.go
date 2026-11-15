package main

import "fmt"

func leaf(v string) *Node { return &Node{Val: v} }

func main() {
	bst := &Node{"4", &Node{"2", leaf("1"), leaf("3")}, &Node{"6", leaf("5"), leaf("7")}}
	fmt.Println("BST 4 2 6 1 3 5 7")
	fmt.Print("  pre:  ")
	PreOrder(bst)
	fmt.Print("\n  in:   ")
	InOrder(bst)
	fmt.Print("  <- sorted\n  post: ")
	PostOrder(bst)
	fmt.Println()

	abcd := &Node{"A", &Node{"B", leaf("D"), nil}, leaf("C")}
	fmt.Println("\n    A\n   / \\\n  B   C\n /\nD")
	fmt.Print("  pre:  ")
	PreOrder(abcd)
	fmt.Print("\n  in:   ")
	InOrder(abcd)
	fmt.Print("\n  post: ")
	PostOrder(abcd)
	fmt.Println()
}
