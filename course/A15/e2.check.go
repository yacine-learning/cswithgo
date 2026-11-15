package main

import "fmt"

func main() {
	known := []string{"break", "case", "const", "continue", "func", "go", "goto", "return", "var"}
	try("a swapped pair of letters", func() string {
		if s, ok := Suggest("fucn", known); !ok || s != "func" {
			return fmt.Sprintf(`Suggest("fucn") = %q, %v; expected "func", true`, s, ok)
		}
		return ""
	})
	try("the closest wins", func() string {
		if s, ok := Suggest("gotoo", known); !ok || s != "goto" {
			return fmt.Sprintf(`Suggest("gotoo") = %q, %v; expected "goto" (1 change), not "go" (3)`, s, ok)
		}
		return ""
	})
	try("too far is no suggestion", func() string {
		if s, ok := Suggest("xyz", known); ok {
			return fmt.Sprintf(`Suggest("xyz") = %q, true; nothing is within 2 changes`, s)
		}
		return ""
	})
	try("a known word needs no suggestion", func() string {
		if s, ok := Suggest("goto", known); ok {
			return fmt.Sprintf(`Suggest("goto") = %q, true; goto is itself known`, s)
		}
		return ""
	})
	try("a tie goes to alphabetical order", func() string {
		for _, ks := range [][]string{{"cat", "bat", "hat"}, {"hat", "cat", "bat"}} {
			if s, ok := Suggest("mat", ks); !ok || s != "bat" {
				return fmt.Sprintf(`Suggest("mat", %v) = %q, %v; all are 1 change away, expected "bat"`, ks, s, ok)
			}
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
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
