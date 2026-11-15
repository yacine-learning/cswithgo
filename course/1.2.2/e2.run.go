package main

import "fmt"

func main() {
	m := NewHashMap()
	m.Put("web-1", 80)
	m.Put("db-1", 5432)
	m.Put("web-1", 8080)
	v, ok := m.Get("web-1")
	fmt.Println("Get(web-1) =", v, ok, " (expected 8080 true)")
	v, ok = m.Get("db-1")
	fmt.Println("Get(db-1)  =", v, ok, " (expected 5432 true)")
	_, ok = m.Get("nope")
	fmt.Println("Get(nope) found:", ok, " (expected false)")
	fmt.Println("Len =", m.Len(), " (expected 2)")
}
