package main

import "fmt"

func expect(progs []Process, q int, want string) func() string {
	return func() string {
		got := fmt.Sprint(RoundRobin(progs, q))
		if got != want {
			return fmt.Sprintf("RoundRobin(%v, %d) = %s, expected %s", progs, q, got, want)
		}
		return ""
	}
}

func main() {
	progs := []Process{{"edit", 5}, {"listf", 1}, {"mad", 4}, {"fap", 2}}
	try("quantum 2", expect(progs, 2, "[listf fap mad edit]"))
	try("a quantum bigger than every program: plain first come, first served", expect(progs, 10, "[edit listf mad fap]"))
	try("quantum 1", expect(progs, 1, "[listf fap mad edit]"))
	try("a program that needs exactly one quantum finishes in its turn", expect([]Process{{"a", 3}, {"b", 3}}, 3, "[a b]"))
	try("no programs", expect(nil, 2, "[]"))
	try("the caller's slice is not changed", func() string {
		js := []Process{{"a", 2}, {"b", 1}}
		RoundRobin(js, 1)
		if js[0].Time != 2 || js[1].Time != 1 {
			return fmt.Sprintf("after the call, the processes passed in became %v", js)
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
