package main

import "sort"

// Link is a possible cable between machines A and B, and what it would cost.
type Link struct{ A, B, Cost int }

// Cheapest connects n machines (0 to n-1) for the least total cost, choosing among links.
// It returns the total cost and the chosen links, or ok = false when some machine
// cannot be reached at all.
//
// Kruskal, 1956: take the links from the cheapest up; keep a link when it joins two
// groups that are not connected yet, skip it when they already are (it would close a loop).
func Cheapest(n int, links []Link) (cost int, chosen []Link, ok bool) {
	sorted := append([]Link(nil), links...)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].Cost < sorted[j].Cost })
	d := NewDSU(n)
	_ = d // Go refuses unused variables: delete this line once your code uses d

	// hole 1: walk the sorted links; keep the ones whose Union succeeds (cost, chosen)

	// hole 2: everything is connected when n-1 links were chosen (a tree on n nodes)
	return cost, chosen, false
}

// The union-find from exercise 1, complete.

// DSU (disjoint-set union) keeps n elements, 0 to n-1, in groups. Each group is a tree:
// every element points to a parent, and the root of the tree names the group.
type DSU struct {
	parent []int
	size   []int // size[r]: number of elements in the group whose root is r
}

// NewDSU starts with n groups of one element: everyone is their own root.
func NewDSU(n int) *DSU {
	d := &DSU{parent: make([]int, n), size: make([]int, n)}
	for i := range d.parent {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

// Find returns the root of x's group. On the way back, it points every element it
// went through straight at the root (path compression), so the next Find is short.
func (d *DSU) Find(x int) int {
	if d.parent[x] == x {
		return x
	}
	root := d.Find(d.parent[x])
	d.parent[x] = root
	return root
}

// Union merges the groups of a and b. It returns false when they were already one group.
// The smaller tree goes under the root of the bigger one (union by size), so trees stay flat.
func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

// Same reports whether a and b are in the same group.
func (d *DSU) Same(a, b int) bool { return d.Find(a) == d.Find(b) }
