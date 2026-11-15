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
	// where each function's code begins in memory (a made-up program)
	funcs := map[int]string{
		0x401000: "runtime.rt0", 0x401200: "main.main", 0x401480: "main.parse", 0x401700: "main.lex",
		0x401a00: "main.eval", 0x401e00: "strconv.Atoi", 0x402000: "fmt.Println", 0x402600: "runtime.mallocgc",
	}
	var tree *Node
	for _, start := range []int{0x401700, 0x401200, 0x402000, 0x401000, 0x401480, 0x401a00, 0x402600, 0x401e00} {
		tree = insert(tree, start)
	}
	// addresses the processor was executing when a profiler sampled it
	for _, pc := range []int{0x401234, 0x4017ff, 0x401a00, 0x401f10, 0x402700, 0x400f00} {
		k, ok := Floor(tree, pc)
		if ok {
			fmt.Printf("%#x is in %s (begins at %#x)\n", pc, funcs[k], k)
		} else {
			fmt.Printf("%#x is before the program's code\n", pc)
		}
	}
}
