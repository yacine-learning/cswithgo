package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	count := 0
	less := func(x, y int) bool { count++; return x < y }
	try("small cases", func() string {
		a := []int{9, 1, 8, 2, 7, 3}
		sorted := []int{1, 2, 3, 7, 8, 9}
		for k := range a {
			if got := Select(a, k, less); got != sorted[k] {
				return fmt.Sprintf("Select(%v, %d) = %d, expected %d", a, k, got, sorted[k])
			}
		}
		if fmt.Sprint(a) != "[9 1 8 2 7 3]" {
			return fmt.Sprintf("Select changed the input to %v", a)
		}
		return ""
	})
	try("duplicates", func() string {
		a := []int{5, 5, 1, 5, 1}
		for k, want := range []int{1, 1, 5, 5, 5} {
			if got := Select(a, k, less); got != want {
				return fmt.Sprintf("Select(%v, %d) = %d, expected %d", a, k, got, want)
			}
		}
		return ""
	})
	try("one value", func() string {
		if got := Select([]int{4}, 0, less); got != 4 {
			return fmt.Sprintf("Select([4], 0) = %d", got)
		}
		return ""
	})
	try("random input: correct, and fewer comparisons than sorting", func() string {
		rng := rand.New(rand.NewSource(9))
		a := make([]int, 100000)
		for i := range a {
			a[i] = rng.Intn(1000000)
		}
		sorted := append([]int(nil), a...)
		sort.Ints(sorted)
		for _, k := range []int{0, 49999, 98999, 99999} {
			count = 0
			if got := Select(a, k, less); got != sorted[k] {
				return fmt.Sprintf("Select(…, %d) = %d, expected %d", k, got, sorted[k])
			}
			if count > 800000 {
				return fmt.Sprintf("%d comparisons for 100 000 values (k=%d); a full sort needs about 1 700 000, quickselect about 300 000", count, k)
			}
		}
		return ""
	})
	try("sorted input does not make it quadratic", func() string {
		a := make([]int, 100000)
		for i := range a {
			a[i] = i
		}
		count = 0
		if got := Select(a, 99000, less); got != 99000 {
			return fmt.Sprintf("Select(0..99999, 99000) = %d", got)
		}
		if count > 800000 {
			return fmt.Sprintf("%d comparisons on sorted input: is the pivot chosen at random?", count)
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
