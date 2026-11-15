package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func valid(h MinHeap) string {
	for i := 1; i < len(h); i++ {
		if h[i] < h[(i-1)/2] {
			return fmt.Sprintf("h[%d]=%d is smaller than its parent h[%d]=%d", i, h[i], (i-1)/2, h[(i-1)/2])
		}
	}
	return ""
}

func main() {
	try("the minimum is on top after each push", func() string {
		var h MinHeap
		least := 1 << 62
		for _, p := range []int{5, 3, 8, 1, 9, 2, 7} {
			h.Push(p)
			least = min(least, p)
			if h[0] != least {
				return fmt.Sprintf("after Push(%d): h[0] = %d, expected %d; heap %v", p, h[0], least, []int(h))
			}
			if msg := valid(h); msg != "" {
				return "after Push: " + msg
			}
		}
		return ""
	})
	try("pops come out in order", func() string {
		var h MinHeap
		for _, p := range []int{5, 3, 8, 1, 9, 2, 7, 3} {
			h.Push(p)
		}
		var got []int
		for len(h) > 0 {
			got = append(got, h.Pop())
			if msg := valid(h); msg != "" {
				return "after Pop: " + msg
			}
		}
		if fmt.Sprint(got) != "[1 2 3 3 5 7 8 9]" {
			return fmt.Sprintf("popped %v, expected [1 2 3 3 5 7 8 9]", got)
		}
		return ""
	})
	try("a node with only a left child", func() string {
		h := MinHeap{1, 5, 3, 6}
		h.Pop()
		if fmt.Sprint([]int(h)) != "[3 5 6]" {
			return fmt.Sprintf("Pop from [1 5 3 6] left %v, expected [3 5 6]", []int(h))
		}
		return ""
	})
	try("10 000 random values in and out", func() string {
		rng := rand.New(rand.NewSource(4))
		var h MinHeap
		var all []int
		for i := 0; i < 10000; i++ {
			v := rng.Intn(1000)
			h.Push(v)
			all = append(all, v)
		}
		sort.Ints(all)
		for i, want := range all {
			if got := h.Pop(); got != want {
				return fmt.Sprintf("pop number %d gave %d, expected %d", i, got, want)
			}
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
