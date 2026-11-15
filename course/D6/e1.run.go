package main

import "fmt"

func keys(c *LRU) []string {
	var out []string
	for el := c.order.Front(); el != nil; el = el.Next() {
		out = append(out, el.Value.(entry).key)
	}
	return out
}

func main() {
	c := NewLRU(3)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	fmt.Println("put a, b, c:     ", keys(c))
	c.Get("a")
	fmt.Println("get a:           ", keys(c))
	c.Put("d", 4)
	fmt.Println("put d (evicts b):", keys(c))
	_, ok := c.Get("b")
	fmt.Println("get b found:", ok)
}
