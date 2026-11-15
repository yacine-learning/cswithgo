package main

import "fmt"

func expect(ids []string, wantStart, wantLen int) func() string {
	return func() string {
		s, n := LongestUnique(ids)
		if s != wantStart || n != wantLen {
			return fmt.Sprintf("LongestUnique(%v) = %d, %d; expected %d, %d", ids, s, n, wantStart, wantLen)
		}
		return ""
	}
}

func main() {
	try("a repeat in the middle", expect([]string{"u1", "u2", "u1", "u3", "u4", "u5", "u3", "u6"}, 1, 5))
	try("no repeat at all", expect([]string{"a", "b", "c"}, 0, 3))
	try("all the same", expect([]string{"x", "x", "x"}, 0, 1))
	try("a repeat that is already outside the window", expect([]string{"a", "b", "b", "a", "c"}, 2, 3))
	try("the window moves past each repeat", expect([]string{"a", "b", "a", "c", "b", "d"}, 2, 4))
	try("a tie: the earliest run", expect([]string{"a", "b", "b", "c"}, 0, 2))
	try("empty", expect(nil, 0, 0))
	try("a long stream", func() string {
		ids := make([]string, 0, 200000)
		for i := 0; i < 200000; i++ {
			ids = append(ids, fmt.Sprint(i%1000))
		}
		s, n := LongestUnique(ids)
		if n != 1000 || s != 0 {
			return fmt.Sprintf("got %d, %d; expected 0, 1000", s, n)
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
