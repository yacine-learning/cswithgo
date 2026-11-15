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
	for i, h := range []string{"web-1", "web-2", "db-1", "cache-1", "web-3", "db-2"} {
		m.Put(h, 10+i)
		fmt.Printf("hash(%q) %% %d = %d\n", h, len(m.buckets), hash(h)%uint32(len(m.buckets)))
	}
	fmt.Println("\n6 entries in 4 buckets:")
	show(m)
	m.Put("queue-1", 20)
	m.Put("queue-2", 21)
	m.Put("lb-1", 22)
	fmt.Printf("\nafter 3 more (9 entries > 2 per bucket): the map grew to %d buckets\n", len(m.buckets))
	show(m)
	v, ok := m.Get("db-1")
	fmt.Println("\nGet(db-1) =", v, ok)
	_, ok = m.Get("db-9")
	fmt.Println("Get(db-9) found:", ok)
}
