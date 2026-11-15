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
	for _, j := range []string{"job1", "job2", "job3"} {
		r.Enqueue(j)
		show(r, "enqueue "+j)
	}
	x, _ := r.Dequeue()
	show(r, "dequeue "+x)
	x, _ = r.Dequeue()
	show(r, "dequeue "+x)
	for _, j := range []string{"job4", "job5", "job6"} {
		r.Enqueue(j)
		show(r, "enqueue "+j)
	}
	fmt.Println("enqueue job7 when full:", r.Enqueue("job7"))
	for r.Len() > 0 {
		x, _ = r.Dequeue()
		show(r, "dequeue "+x)
	}
}
