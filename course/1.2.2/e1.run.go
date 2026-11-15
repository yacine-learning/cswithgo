package main

import "fmt"

func show(m *HashMap) {
	for i, b := range m.buckets {
		fmt.Printf("  bucket %d:", i)
		for _, e := range b {
			fmt.Printf(" %s=%d", e.key, e.value)
		}
		fmt.Println()
	}
}

func main() {
	m := NewHashMap()
	for i, h := range []string{"start", "loop", "sum", "count", "next", "done"} {
		m.Put(h, []int{100, 104, 200, 201, 112, 120}[i]) // memory addresses
		fmt.Printf("hash(%q) %% %d = %d\n", h, len(m.buckets), hash(h)%uint32(len(m.buckets)))
	}
	fmt.Println("\n6 entries in 4 buckets:")
	show(m)
	m.Put("print", 300)
	m.Put("buffer", 202)
	m.Put("exit", 124)
	fmt.Printf("\nafter 3 more (9 entries > 2 per bucket): the map grew to %d buckets\n", len(m.buckets))
	show(m)
	v, ok := m.Get("count")
	fmt.Println("\nGet(count) =", v, ok)
	_, ok = m.Get("missing")
	fmt.Println("Get(missing) found:", ok)
}
