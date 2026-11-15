package main

import "fmt"

func expect(a, b []Entry, want string) func() string {
	return func() string {
		if got := fmt.Sprint(Merge(a, b)); got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	}
}

func main() {
	a := []Entry{{100, "a1"}, {105, "a2"}, {112, "a3"}, {130, "a4"}}
	b := []Entry{{101, "b1"}, {105, "b2"}, {140, "b3"}}
	try("interleaved, equal times: a first", expect(a, b, "[{100 a1} {101 b1} {105 a2} {105 b2} {112 a3} {130 a4} {140 b3}]"))
	try("b entirely before a", expect([]Entry{{5, "a"}}, []Entry{{1, "b1"}, {2, "b2"}}, "[{1 b1} {2 b2} {5 a}]"))
	try("one log empty", expect(nil, []Entry{{1, "b"}}, "[{1 b}]"))
	try("both empty", expect(nil, nil, "[]"))
	try("the inputs are not changed", func() string {
		x := []Entry{{1, "x"}, {3, "y"}}
		y := []Entry{{2, "z"}}
		Merge(x, y)
		if fmt.Sprint(x, y) != "[{1 x} {3 y}] [{2 z}]" {
			return fmt.Sprintf("after Merge the inputs are %v %v", x, y)
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
