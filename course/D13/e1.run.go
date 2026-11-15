package main

import "fmt"

func main() {
	d := NewDSU(8)
	links := [][2]int{{0, 1}, {2, 3}, {1, 3}, {4, 5}, {0, 2}, {6, 7}, {5, 7}}
	for _, l := range links {
		merged := d.Union(l[0], l[1])
		fmt.Printf("link %d-%d  merged: %-5v  parents: %v\n", l[0], l[1], merged, d.parent)
	}
	fmt.Println()
	for _, q := range [][2]int{{0, 3}, {1, 2}, {3, 4}, {4, 7}} {
		fmt.Printf("%d and %d in the same group: %v\n", q[0], q[1], d.Same(q[0], q[1]))
	}

	// a long chain of links: how deep do the trees get?
	n := 1 << 16
	c := NewDSU(n)
	for i := 1; i < n; i++ {
		c.Union(i-1, i)
	}
	depth := 0
	for x := 0; c.parent[x] != x; x = c.parent[x] {
		depth++
	}
	fmt.Printf("\n%d elements linked in a chain, one at a time: element 0 is %d steps from its root\n", n, depth)
}
