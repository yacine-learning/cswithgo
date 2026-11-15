package main

import "fmt"

func expect(in []int, want string) func() string {
	return func() string {
		a := append([]int(nil), in...)
		k := Dedup(a)
		if k < 0 || k > len(a) {
			return fmt.Sprintf("Dedup(%v) returned %d", in, k)
		}
		if got := fmt.Sprint(a[:k]); got != want {
			return fmt.Sprintf("Dedup(%v): a[:k] = %s, expected %s", in, got, want)
		}
		return ""
	}
}

func main() {
	try("repeats in the middle and at the end", expect([]int{1, 1, 2, 3, 3, 3, 7, 9, 9}, "[1 2 3 7 9]"))
	try("no repeats", expect([]int{1, 2, 3}, "[1 2 3]"))
	try("all the same", expect([]int{4, 4, 4, 4}, "[4]"))
	try("one value", expect([]int{5}, "[5]"))
	try("empty", expect(nil, "[]"))
	try("negative values", expect([]int{-3, -3, -1, 0, 0}, "[-3 -1 0]"))
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
