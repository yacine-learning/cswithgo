package main

import "fmt"

func main() {
	try("reverse five values", func() string {
		l := FromValues(1, 2, 3, 4, 5)
		l.Reverse()
		if got := fmt.Sprint(l.Values()); got != "[5 4 3 2 1]" {
			return "reversing [1 2 3 4 5] gave " + got
		}
		return ""
	})
	try("an empty list", func() string {
		l := FromValues()
		l.Reverse()
		if len(l.Values()) != 0 {
			return fmt.Sprintf("reversing an empty list gave %v", l.Values())
		}
		return ""
	})
	try("one value", func() string {
		l := FromValues(42)
		l.Reverse()
		if got := fmt.Sprint(l.Values()); got != "[42]" {
			return "reversing [42] gave " + got
		}
		return ""
	})
	try("reversing twice gives the original", func() string {
		l := FromValues(3, 1, 4, 1, 5, 9)
		l.Reverse()
		l.Reverse()
		if got := fmt.Sprint(l.Values()); got != "[3 1 4 1 5 9]" {
			return "reversing [3 1 4 1 5 9] twice gave " + got
		}
		return ""
	})
	try("the same nodes are reused, not copied", func() string {
		l := FromValues(1, 2, 3)
		first := l.head.next
		l.Reverse()
		last := l.head.next
		for last.next != nil {
			last = last.next
		}
		if last != first {
			return "the old first node is not the new last node: were new nodes created?"
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
