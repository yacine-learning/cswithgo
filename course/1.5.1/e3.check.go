package main

import (
	"fmt"
	"math/bits"
)

func run(n, first int) (got, calls int) {
	got = FirstTrue(n, func(b int) bool {
		calls++
		if b < 0 || b >= n {
			panic(fmt.Sprintf("holds(%d) called, but the candidates are 0..%d", b, n-1))
		}
		return b >= first
	})
	return got, calls
}

func main() {
	for _, c := range []struct {
		name     string
		n, first int
	}{
		{"in the middle", 1000, 437},
		{"true from the start", 1000, 0},
		{"true only for the last candidate", 1000, 999},
		{"never true", 1000, 1000},
		{"one candidate, true", 1, 0},
		{"one candidate, false", 1, 1},
		{"no candidates", 0, 0},
	} {
		try(c.name, func() string {
			if got, _ := run(c.n, c.first); got != c.first {
				return fmt.Sprintf("FirstTrue(%d) = %d, expected %d", c.n, got, c.first)
			}
			return ""
		})
	}
	try("about log2(n) calls to holds", func() string {
		n := 1_000_000
		limit := bits.Len(uint(n)) + 1 // 21
		for _, first := range []int{0, 1, 123_456, 999_999, n} {
			if got, calls := run(n, first); got != first || calls > limit {
				return fmt.Sprintf("for 1 000 000 candidates: %d calls to holds (answer %d); at most %d are needed", calls, got, limit)
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
