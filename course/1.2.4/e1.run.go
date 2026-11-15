package main

import "fmt"

func show(r *Ring, what string) {
	fmt.Printf("%-16s head=%d count=%d  [", what, r.head, r.count)
	for i, s := range r.buf {
		if s == "" {
			s = "·"
		}
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%-5s", s)
	}
	fmt.Println("]")
}

func main() {
	r := NewRing(4)
	show(r, "empty")
	for _, j := range []string{"a", "b", "c"} {
		r.Enqueue(j)
		show(r, "enqueue "+j)
	}
	x, _ := r.Dequeue()
	show(r, "dequeue "+x)
	x, _ = r.Dequeue()
	show(r, "dequeue "+x)
	for _, j := range []string{"d", "e", "f"} {
		r.Enqueue(j)
		show(r, "enqueue "+j)
	}
	fmt.Println("enqueue g when full:", r.Enqueue("g"))
	for r.Len() > 0 {
		x, _ = r.Dequeue()
		show(r, "dequeue "+x)
	}
}
