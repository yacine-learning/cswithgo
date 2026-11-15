package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	counts := map[string]int{
		"connection refused": 412, "timeout": 977, "401 unauthorized": 58,
		"disk full": 3, "502 bad gateway": 412, "DNS lookup failed": 120,
	}
	expect := func(k int, want string) func() string {
		return func() string {
			for i := 0; i < 20; i++ {
				if got := fmt.Sprint(TopK(counts, k)); got != want {
					return fmt.Sprintf("TopK(counts, %d) = %s, expected %s", k, got, want)
				}
			}
			return ""
		}
	}
	try("top 1", expect(1, "[timeout]"))
	try("top 3, with a tie", expect(3, "[timeout 502 bad gateway connection refused]"))
	try("k larger than the number of messages", expect(10, "[timeout 502 bad gateway connection refused DNS lookup failed 401 unauthorized disk full]"))
	try("k = 0", expect(0, "[]"))
	try("a tie at the cut", func() string {
		c := map[string]int{"b": 5, "a": 5, "c": 5, "d": 1}
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(TopK(c, 2)); got != "[a b]" {
				return "TopK({a:5 b:5 c:5 d:1}, 2) = " + got + ", expected [a b]"
			}
		}
		return ""
	})
	try("10 000 messages", func() string {
		rng := rand.New(rand.NewSource(8))
		c := map[string]int{}
		for i := 0; i < 10000; i++ {
			c[fmt.Sprintf("err-%05d", i)] = rng.Intn(500)
		}
		all := make([]string, 0, len(c))
		for m := range c {
			all = append(all, m)
		}
		sort.Slice(all, func(i, j int) bool {
			if c[all[i]] != c[all[j]] {
				return c[all[i]] > c[all[j]]
			}
			return all[i] < all[j]
		})
		if got, want := fmt.Sprint(TopK(c, 25)), fmt.Sprint(all[:25]); got != want {
			return "the top 25 differ from a full sort"
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
