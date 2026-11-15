package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var root *Node
	for _, k := range []int{4, 6, 7, 5, 2, 1, 3} {
		root = Insert(root, k)
	}
	fmt.Println("        4")
	fmt.Println("      /   \\")
	fmt.Println("     2     6")
	fmt.Println("    / \\   / \\")
	fmt.Println("   1   3 5   7")
	for _, k := range []int{5, 1, 10} {
		found, v := Search(root, k)
		fmt.Printf("search %-2d found=%-5v visited %d of 7 nodes\n", k, found, v)
	}

	rng := rand.New(rand.NewSource(1))
	var random, sorted *Node
	for _, k := range rng.Perm(1000) {
		random = Insert(random, k)
	}
	for k := 0; k < 1000; k++ {
		sorted = Insert(sorted, k)
	}
	_, vr := Search(random, 999)
	_, vs := Search(sorted, 999)
	fmt.Printf("\n1 000 keys in random order: height %d, search 999 visits %d nodes\n", Height(random), vr)
	fmt.Printf("1 000 keys in sorted order: height %d, search 999 visits %d nodes  (a list!)\n", Height(sorted), vs)
}
