package main

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
	// hole 1: x is a root when it is its own parent
	// hole 2: otherwise find the root of its parent, point x at that root, return it
	return x
}

// Union merges the groups of a and b. It returns false when they were already one group.
// The smaller tree goes under the root of the bigger one (union by size), so trees stay flat.
func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	// hole 3: same root, nothing to do
	// hole 4: make ra the root of the bigger group (swap if needed), then attach rb under ra
	//         and add its size
	_, _ = ra, rb // Go refuses unused variables: delete this line once your code uses ra and rb
	return true
}

// Same reports whether a and b are in the same group.
func (d *DSU) Same(a, b int) bool { return d.Find(a) == d.Find(b) }
