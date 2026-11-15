package main

import "fmt"

func main() {
	try("first in, first out", func() string {
		r := NewRing(4)
		r.Enqueue("a")
		r.Enqueue("b")
		r.Enqueue("c")
		var got []string
		for r.Len() > 0 {
			x, _ := r.Dequeue()
			got = append(got, x)
		}
		if fmt.Sprint(got) != "[a b c]" {
			return fmt.Sprintf("enqueue a, b, c then dequeue all gave %v", got)
		}
		return ""
	})
	try("a full queue refuses", func() string {
		r := NewRing(2)
		if !r.Enqueue("a") || !r.Enqueue("b") {
			return "enqueueing 2 items into a queue of capacity 2 was refused"
		}
		if r.Enqueue("c") {
			return "a third item was accepted into a queue of capacity 2"
		}
		return ""
	})
	try("an empty queue refuses", func() string {
		r := NewRing(2)
		if _, ok := r.Dequeue(); ok {
			return "Dequeue on an empty queue returned ok=true"
		}
		return ""
	})
	try("wrapping around many times", func() string {
		r := NewRing(3)
		next, want := 0, 0
		for round := 0; round < 100; round++ {
			for k := 0; k < 4 && r.Enqueue(fmt.Sprint(next)); k++ { // at most 4 tries: the capacity is 3
				next++
			}
			for i := 0; i < 2; i++ {
				x, ok := r.Dequeue()
				if !ok || x != fmt.Sprint(want) {
					return fmt.Sprintf("round %d: dequeued %q (ok=%v), expected %q", round, x, ok, fmt.Sprint(want))
				}
				want++
			}
		}
		return ""
	})
	try("Len follows enqueues and dequeues", func() string {
		r := NewRing(5)
		r.Enqueue("a")
		r.Enqueue("b")
		r.Dequeue()
		r.Enqueue("c")
		if r.Len() != 2 {
			return fmt.Sprintf("Len() = %d after 3 enqueues and 1 dequeue, expected 2", r.Len())
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
