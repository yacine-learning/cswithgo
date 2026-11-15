package main

import "fmt"

func expect(limit, window int, times []int, want string) func() string {
	return func() string {
		l := &Limiter{Limit: limit, Window: window}
		var got []bool
		for _, t := range times {
			got = append(got, l.Allow(t))
		}
		if fmt.Sprint(got) != want {
			return fmt.Sprintf("limit %d per %d s, requests at %v:\n      allowed %v\n      expected %s", limit, window, times, got, want)
		}
		return ""
	}
}

func main() {
	try("a burst is cut at the limit", expect(3, 10, []int{0, 1, 2, 3, 4}, "[true true true false false]"))
	try("a request leaves the window after Window seconds", expect(3, 10, []int{0, 1, 2, 9, 10, 11, 12}, "[true true true false true true true]"))
	try("refused requests do not count", expect(2, 10, []int{0, 0, 0, 0, 10, 10, 10}, "[true true false false true true false]"))
	try("a long quiet gap empties the window", expect(2, 5, []int{0, 1, 100, 100, 100}, "[true true true true false]"))
	try("the queue does not grow without bound", func() string {
		l := &Limiter{Limit: 5, Window: 10}
		for t := 0; t < 100000; t++ {
			l.Allow(t)
		}
		if len(l.times) > 5 {
			return fmt.Sprintf("after 100 000 requests the limiter holds %d times; it never needs more than Limit", len(l.times))
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
