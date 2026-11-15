package main

import "container/list"

type entry struct {
	key   string
	value int
}

// LRU keeps the capacity most recently used keys.
type LRU struct {
	capacity int
	order    *list.List               // front: most recently used
	items    map[string]*list.Element // key -> its element in order
}

func NewLRU(capacity int) *LRU {
	return &LRU{capacity: capacity, order: list.New(), items: map[string]*list.Element{}}
}

// Get returns the value of key and marks it as just used.
func (c *LRU) Get(key string) (int, bool) {
	el, ok := c.items[key]
	if !ok {
		return 0, false
	}
	c.order.MoveToFront(el)
	return el.Value.(entry).value, true
}

// Put sets the value of key and marks it as just used. If the cache is then over
// capacity, the least recently used key is evicted.
func (c *LRU) Put(key string, value int) {
	if el, ok := c.items[key]; ok {
		el.Value = entry{key, value}
		c.order.MoveToFront(el)
		return
	}
	c.items[key] = c.order.PushFront(entry{key, value})
	if c.order.Len() > c.capacity {
		back := c.order.Back()
		c.order.Remove(back)
		delete(c.items, back.Value.(entry).key)
	}
}
