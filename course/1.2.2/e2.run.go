package main

import "fmt"

func main() {
	m := NewHashMap()
	m.Put("hydrogen", 1)
	m.Put("helium", 2)
	m.Put("hydrogen", 100)
	v, ok := m.Get("hydrogen")
	fmt.Println("Get(hydrogen) =", v, ok, " (expected 100 true)")
	v, ok = m.Get("helium")
	fmt.Println("Get(helium)  =", v, ok, " (expected 2 true)")
	_, ok = m.Get("nope")
	fmt.Println("Get(nope) found:", ok, " (expected false)")
	fmt.Println("Len =", m.Len(), " (expected 2)")
}
