package main

import "fmt"

func main() {
	try("everyone starts alone", func() string {
		d := NewDSU(4)
		for i := 0; i < 4; i++ {
			if d.Find(i) != i {
				return fmt.Sprintf("in a new DSU, Find(%d) = %d", i, d.Find(i))
			}
		}
		if d.Same(0, 1) {
			return "in a new DSU, 0 and 1 are in the same group"
		}
		return ""
	})
	try("union joins two groups", func() string {
		d := NewDSU(4)
		if !d.Union(0, 1) {
			return "Union(0, 1) on separate elements returned false"
		}
		if !d.Same(0, 1) {
			return "after Union(0, 1), 0 and 1 are not in the same group"
		}
		if d.Same(0, 2) {
			return "after Union(0, 1), 0 and 2 are in the same group"
		}
		return ""
	})
	try("groups join through their roots", func() string {
		d := NewDSU(6)
		d.Union(0, 1)
		d.Union(2, 3)
		d.Union(1, 3) // must join the groups {0,1} and {2,3}, not just the elements 1 and 3
		for _, p := range [][2]int{{0, 3}, {1, 2}, {0, 2}} {
			if !d.Same(p[0], p[1]) {
				return fmt.Sprintf("after Union(0,1), Union(2,3), Union(1,3): %d and %d are not in the same group", p[0], p[1])
			}
		}
		if d.Same(1, 4) {
			return "4 was never linked, but it is in 1's group"
		}
		return ""
	})
	try("union of one group returns false", func() string {
		d := NewDSU(3)
		d.Union(0, 1)
		d.Union(1, 2)
		if d.Union(0, 2) {
			return "after Union(0,1) and Union(1,2), Union(0, 2) returned true: they were already one group"
		}
		return ""
	})
	try("sizes are kept", func() string {
		d := NewDSU(5)
		d.Union(0, 1)
		d.Union(2, 3)
		d.Union(3, 4)
		d.Union(1, 4)
		if s := d.size[d.Find(0)]; s != 5 {
			return fmt.Sprintf("all 5 elements in one group, but the size at the root is %d", s)
		}
		return ""
	})
	// last: too slow without path compression or union by size, the page stops it
	try("a million operations stay fast", func() string {
		n := 200000
		d := NewDSU(n)
		for i := 1; i < n; i++ {
			d.Union(i, i-1)
		}
		for i := 0; i < n; i++ {
			if d.Find(i) != d.Find(0) {
				return fmt.Sprintf("after linking a chain of %d, %d is not in 0's group", n, i)
			}
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
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
