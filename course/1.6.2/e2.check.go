package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	try("one partition", func() string {
		a := []int{7, 2, 9, 4, 1, 8, 5}
		p := partition(a, 0, len(a)-1)
		if p != 3 || a[3] != 5 {
			return fmt.Sprintf("partition(%v) returned %d and left %v; the pivot 5 belongs at index 3", []int{7, 2, 9, 4, 1, 8, 5}, p, a)
		}
		for i := 0; i < 3; i++ {
			if a[i] >= 5 {
				return fmt.Sprintf("after partition, %v: a[%d]=%d is before the pivot but not smaller", a, i, a[i])
			}
		}
		for i := 4; i < 7; i++ {
			if a[i] < 5 {
				return fmt.Sprintf("after partition, %v: a[%d]=%d is after the pivot but smaller", a, i, a[i])
			}
		}
		return ""
	})
	try("partition of a sub-range leaves the rest alone", func() string {
		a := []int{100, 3, 1, 2, -100}
		p := partition(a, 1, 3)
		if a[0] != 100 || a[4] != -100 || p != 2 || a[2] != 2 {
			return fmt.Sprintf("partition(a, 1, 3) on [100 3 1 2 -100] gave %v, pivot at %d", a, p)
		}
		return ""
	})
	try("sorts with duplicates", func() string {
		a := []int{3, 1, 3, 2, 3, 1, 2}
		QuickSort(a, 0, len(a)-1)
		if fmt.Sprint(a) != "[1 1 2 2 3 3 3]" {
			return fmt.Sprintf("got %v", a)
		}
		return ""
	})
	try("sorts 20 000 random values", func() string {
		rng := rand.New(rand.NewSource(3))
		a := make([]int, 20000)
		for i := range a {
			a[i] = rng.Intn(1000)
		}
		QuickSort(a, 0, len(a)-1)
		if !sort.IntsAreSorted(a) {
			return "the result is not sorted"
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
