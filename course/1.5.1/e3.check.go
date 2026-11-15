package main

import (
	"fmt"
	"math/bits"
)

func run(n, first int) (got, calls int) {
	got = FirstBad(n, func(b int) bool {
		calls++
		if b < 0 || b >= n {
			panic(fmt.Sprintf("isBad(%d) called, but builds are 0..%d", b, n-1))
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
		{"the very first build is bad", 1000, 0},
		{"only the last build is bad", 1000, 999},
		{"no build is bad", 1000, 1000},
		{"a single build, bad", 1, 0},
		{"a single build, good", 1, 1},
		{"no builds", 0, 0},
	} {
		try(c.name, func() string {
			if got, _ := run(c.n, c.first); got != c.first {
				return fmt.Sprintf("FirstBad(%d) = %d, expected %d", c.n, got, c.first)
			}
			return ""
		})
	}
	try("about log2(n) calls to isBad", func() string {
		n := 1_000_000
		limit := bits.Len(uint(n)) + 1 // 21
		for _, first := range []int{0, 1, 123_456, 999_999, n} {
			if got, calls := run(n, first); got != first || calls > limit {
				return fmt.Sprintf("for 1 000 000 builds: %d calls to isBad (answer %d); at most %d are needed", calls, got, limit)
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
