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

func main() {
	// start of each period, millions of years ago (International Chronostratigraphic Chart, rounded)
	periods := map[int]string{
		-539: "Cambrian", -485: "Ordovician", -444: "Silurian", -419: "Devonian", -359: "Carboniferous",
		-299: "Permian", -252: "Triassic", -201: "Jurassic", -145: "Cretaceous", -66: "Paleogene",
		-23: "Neogene", -3: "Quaternary",
	}
	var tree *Node
	for _, start := range []int{-252, -419, -66, -539, -359, -201, -23, -485, -444, -299, -145, -3} {
		tree = insert(tree, start)
	}
	for _, t := range []int{-500, -300, -150, -65, -1, -600} {
		k, ok := Floor(tree, t)
		if ok {
			fmt.Printf("%4d million years: %s (began %d)\n", t, periods[k], k)
		} else {
			fmt.Printf("%4d million years: before the Cambrian\n", t)
		}
	}
}
