package main

import "fmt"

func main() {
	try("put then get", func() string {
		c := NewLRU(2)
		c.Put("a", 1)
		if v, ok := c.Get("a"); !ok || v != 1 {
			return fmt.Sprintf("Get(\"a\") = %d, %v after Put(\"a\", 1)", v, ok)
		}
		return ""
	})
	try("evicts the least recently used", func() string {
		c := NewLRU(2)
		c.Put("a", 1)
		c.Put("b", 2)
		c.Put("c", 3)
		if _, ok := c.Get("a"); ok {
			return "after Put a, b, c in a cache of 2, \"a\" is still there"
		}
		if _, ok := c.Get("b"); !ok {
			return "\"b\" was evicted instead of \"a\""
		}
		return ""
	})
	try("Get counts as a use", func() string {
		c := NewLRU(2)
		c.Put("a", 1)
		c.Put("b", 2)
		c.Get("a")
		c.Put("c", 3)
		if _, ok := c.Get("a"); !ok {
			return "Put a, Put b, Get a, Put c: \"a\" was evicted, but \"b\" was the least recently used"
		}
		if _, ok := c.Get("b"); ok {
			return "Put a, Put b, Get a, Put c: \"b\" should have been evicted"
		}
		return ""
	})
	try("putting an existing key updates it and counts as a use", func() string {
		c := NewLRU(2)
		c.Put("a", 1)
		c.Put("b", 2)
		c.Put("a", 10)
		c.Put("c", 3)
		if v, ok := c.Get("a"); !ok || v != 10 {
			return fmt.Sprintf("Get(\"a\") = %d, %v; expected 10, true", v, ok)
		}
		if c.order.Len() != 2 || len(c.items) != 2 {
			return fmt.Sprintf("the cache holds %d elements and %d map entries, expected 2 and 2", c.order.Len(), len(c.items))
		}
		return ""
	})
	try("the map and the list stay the same size", func() string {
		c := NewLRU(3)
		for i := 0; i < 100; i++ {
			c.Put(fmt.Sprint(i%7), i)
			c.Get(fmt.Sprint(i % 5))
		}
		if c.order.Len() != len(c.items) || c.order.Len() > 3 {
			return fmt.Sprintf("list has %d elements, map has %d keys (capacity 3): is the evicted key deleted from the map?", c.order.Len(), len(c.items))
		}
		return ""
	})
}

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
