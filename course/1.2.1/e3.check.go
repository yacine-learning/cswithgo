package main

import "fmt"

func main() {
	try("two paths from the same prefix keep their own last step", func() string {
		prefix := make([]int, 0, 10)
		prefix = append(prefix, 1, 4, 7)
		a := Extend(prefix, 8)
		b := Extend(prefix, 9)
		if len(a) != 4 || a[3] != 8 {
			return fmt.Sprintf("the first path is %v after building the second one; expected [1 4 7 8]", a)
		}
		if len(b) != 4 || b[3] != 9 {
			return fmt.Sprintf("the second path is %v, expected [1 4 7 9]", b)
		}
		return ""
	})
	try("the prefix is unchanged", func() string {
		prefix := []int{1, 4, 7}
		Extend(prefix, 8)
		if fmt.Sprint(prefix) != "[1 4 7]" {
			return fmt.Sprintf("the prefix became %v", prefix)
		}
		return ""
	})
	try("changing a result does not change the prefix", func() string {
		prefix := make([]int, 3, 10)
		prefix[0], prefix[1], prefix[2] = 1, 4, 7
		r := Extend(prefix, 8)
		r[0] = 99
		if prefix[0] != 1 {
			return "setting result[0] also changed prefix[0]: they share one array"
		}
		return ""
	})
	try("an empty prefix", func() string {
		r := Extend(nil, 5)
		if len(r) != 1 || r[0] != 5 {
			return fmt.Sprintf("Extend(nil, 5) = %v", r)
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
