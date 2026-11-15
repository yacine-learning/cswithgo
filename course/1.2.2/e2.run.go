package main

import "fmt"

func main() {
	m := NewHashMap()
	m.Put("loop", 1)
	m.Put("sum", 2)
	m.Put("loop", 100)
	v, ok := m.Get("loop")
	fmt.Println("Get(loop) =", v, ok, " (expected 100 true)")
	v, ok = m.Get("sum")
	fmt.Println("Get(sum)  =", v, ok, " (expected 2 true)")
	_, ok = m.Get("nope")
	fmt.Println("Get(nope) found:", ok, " (expected false)")
	fmt.Println("Len =", m.Len(), " (expected 2)")
}
