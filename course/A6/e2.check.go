package main

import "fmt"

func expect(in []Interval, want string) func() string {
	return func() string {
		if got := fmt.Sprint(MergeIntervals(in)); got != want {
			return fmt.Sprintf("MergeIntervals(%v) = %s, expected %s", in, got, want)
		}
		return ""
	}
}

func main() {
	try("overlapping, touching and separate", expect([]Interval{{120, 180}, {0, 30}, {150, 200}, {30, 45}, {300, 360}, {170, 190}},
		"[{0 45} {120 200} {300 360}]"))
	try("one interval inside another", expect([]Interval{{0, 100}, {10, 20}}, "[{0 100}]"))
	try("already separate and sorted", expect([]Interval{{0, 1}, {5, 6}}, "[{0 1} {5 6}]"))
	try("unsorted input", expect([]Interval{{50, 60}, {0, 10}, {20, 30}}, "[{0 10} {20 30} {50 60}]"))
	try("empty", expect(nil, "[]"))
	try("the caller's slice is not changed", func() string {
		in := []Interval{{5, 9}, {0, 6}}
		MergeIntervals(in)
		if fmt.Sprint(in) != "[{5 9} {0 6}]" {
			return fmt.Sprintf("after the call the input became %v", in)
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
