package main

import "fmt"

func expect(s string, wantOK bool, wantPos int) func() string {
	return func() string {
		ok, pos := Balanced(s)
		if ok != wantOK || pos != wantPos {
			return fmt.Sprintf("Balanced(%q) = %v, %d; expected %v, %d", s, ok, pos, wantOK, wantPos)
		}
		return ""
	}
}

func main() {
	try("a balanced formula", expect(`{a + [1 * (2 - b)]} / {c}`, true, -1))
	try("no brackets at all", expect("plain text", true, -1))
	try("empty string", expect("", true, -1))
	try("wrong kind of closer", expect("{[}]", false, 2))
	try("closer with nothing open", expect("a)b", false, 1))
	try("opener left unclosed: report the innermost", expect("{ a: [1, 2", false, 5))
	try("closed in the wrong order", expect("([)]", false, 2))
	try("deep nesting", expect("((((((((((x))))))))))", true, -1))
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
