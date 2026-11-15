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
	for i, h := range []string{"hydrogen", "helium", "lithium", "carbon", "nitrogen", "oxygen"} {
		m.Put(h, []int{1, 2, 3, 6, 7, 8}[i]) // atomic numbers
		fmt.Printf("hash(%q) %% %d = %d\n", h, len(m.buckets), hash(h)%uint32(len(m.buckets)))
	}
	fmt.Println("\n6 entries in 4 buckets:")
	show(m)
	m.Put("neon", 10)
	m.Put("sodium", 11)
	m.Put("iron", 26)
	fmt.Printf("\nafter 3 more (9 entries > 2 per bucket): the map grew to %d buckets\n", len(m.buckets))
	show(m)
	v, ok := m.Get("carbon")
	fmt.Println("\nGet(carbon) =", v, ok)
	_, ok = m.Get("unobtainium")
	fmt.Println("Get(unobtainium) found:", ok)
}
