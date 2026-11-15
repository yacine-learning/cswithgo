package main

import "fmt"

func main() {
	a := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	expect := func(arr []int, t, want int) func() string {
		return func() string {
			if got := Search(arr, t); got != want {
				return fmt.Sprintf("Search(%v, %d) = %d, expected %d", arr, t, got, want)
			}
			return ""
		}
	}
	try("a value in the middle", expect(a, 23, 5))
	try("the first element", expect(a, 2, 0))
	try("the last element", expect(a, 91, 9))
	try("a missing value", expect(a, 50, -1))
	try("smaller than everything", expect(a, 1, -1))
	try("larger than everything", expect(a, 100, -1))
	try("a single element", expect([]int{7}, 7, 0))
	try("two elements, the second one", expect([]int{3, 9}, 9, 1))
	try("an empty slice", expect(nil, 4, -1))
	try("every element of 1 000 is found", func() string {
		big := make([]int, 1000)
		for i := range big {
			big[i] = i * 3
		}
		for i := range big {
			if got := Search(big, i*3); got != i {
				return fmt.Sprintf("Search(…, %d) = %d, expected %d", i*3, got, i)
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
