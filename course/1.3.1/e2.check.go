package main

import "fmt"

func main() {
	try("small values", func() string {
		for n, want := range []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34} {
			if got := Fib(n, map[int]int{}); got != want {
				return fmt.Sprintf("Fib(%d) = %d, expected %d", n, got, want)
			}
		}
		return ""
	})
	try("Fib(30) in fewer than 100 calls", func() string {
		Calls = 0
		v := Fib(30, map[int]int{})
		if v != 832040 {
			return fmt.Sprintf("Fib(30) = %d, expected 832040", v)
		}
		if Calls > 100 {
			return fmt.Sprintf("Fib(30) took %d calls: is every value computed once?", Calls)
		}
		return ""
	})
	try("Fib(90)", func() string {
		if v := Fib(90, map[int]int{}); v != 2880067194370816120 {
			return fmt.Sprintf("Fib(90) = %d, expected 2880067194370816120", v)
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
