package main

import "fmt"

func main() {
	try("the first push into an empty Vec", func() string {
		var v Vec
		v.Push(7)
		if v.Len() != 1 || v.Get(0) != 7 {
			return fmt.Sprintf("after Push(7): Len()=%d, Get(0)=%d; expected 1 and 7", v.Len(), v.Get(0))
		}
		return ""
	})
	try("1 000 pushes keep every value", func() string {
		var v Vec
		for i := 0; i < 1000; i++ {
			v.Push(i)
		}
		if v.Len() != 1000 {
			return fmt.Sprintf("Len()=%d, expected 1000", v.Len())
		}
		for i := 0; i < 1000; i++ {
			if v.Get(i) != i {
				return fmt.Sprintf("Get(%d)=%d, expected %d: were the old values copied when the array grew?", i, v.Get(i), i)
			}
		}
		return ""
	})
	try("capacity doubles (1, 2, 4, …)", func() string {
		var v Vec
		var caps []int
		for i := 0; i < 9; i++ {
			before := v.Cap()
			v.Push(i)
			if v.Cap() != before {
				caps = append(caps, v.Cap())
			}
		}
		want := []int{1, 2, 4, 8, 16}
		if fmt.Sprint(caps) != fmt.Sprint(want) {
			return fmt.Sprintf("capacities after each growth were %v, expected %v", caps, want)
		}
		return ""
	})
	try("fewer than 2 copies per element (amortized O(1))", func() string {
		var v Vec
		for i := 0; i < 1000; i++ {
			v.Push(i)
		}
		if v.Copies > 2000 {
			return fmt.Sprintf("%d copies for 1 000 pushes", v.Copies)
		}
		if v.Copies == 0 {
			return "Copies is 0: count each element copied"
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
