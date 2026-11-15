package main

import "fmt"

func expect(window int, times []int, want string) func() string {
	return func() string {
		c := &Counter{Window: window}
		var got []int
		for _, t := range times {
			got = append(got, c.Click(t))
		}
		if fmt.Sprint(got) != want {
			return fmt.Sprintf("window %d s, clicks at %v:\n      counts %v\n      expected %s", window, times, got, want)
		}
		return ""
	}
}

func main() {
	try("a burst", expect(10, []int{0, 1, 2, 3, 4}, "[1 2 3 4 5]"))
	try("a click leaves the window after Window seconds", expect(10, []int{0, 1, 2, 9, 10, 11, 12}, "[1 2 3 4 4 4 4]"))
	try("clicks at the same instant", expect(10, []int{5, 5, 5, 15, 15}, "[1 2 3 1 2]"))
	try("a long quiet gap empties the window", expect(5, []int{0, 1, 100, 100, 101}, "[1 2 1 2 3]"))
	try("the queue does not grow without bound", func() string {
		c := &Counter{Window: 10}
		for t := 0; t < 100000; t++ {
			c.Click(t)
		}
		if len(c.times) > 10 {
			return fmt.Sprintf("after 100 000 clicks, one per second, the counter holds %d times; only 10 are in a 10 s window", len(c.times))
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
