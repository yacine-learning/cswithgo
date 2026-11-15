package main

import "fmt"

func main() {
	want := []int{1, 0, 0, 2, 10, 4, 40, 92}
	try("small boards", func() string {
		for i, w := range want[:4] {
			if got := Queens(i + 1); got != w {
				return fmt.Sprintf("Queens(%d) = %d, expected %d", i+1, got, w)
			}
		}
		return ""
	})
	try("up to 8 × 8", func() string {
		for i, w := range want {
			if got := Queens(i + 1); got != w {
				return fmt.Sprintf("Queens(%d) = %d, expected %d", i+1, got, w)
			}
		}
		return ""
	})
	try("10 × 10", func() string {
		if got := Queens(10); got != 724 {
			return fmt.Sprintf("Queens(10) = %d, expected 724", got)
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
