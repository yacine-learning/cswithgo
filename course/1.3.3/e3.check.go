package main

import "fmt"

func main() {
	thesis := &Dir{"thesis", []*Dir{
		{"chapters", []*Dir{{"intro.tex", nil}, {"figures", []*Dir{{"plot.svg", nil}}}}},
		{"data", []*Dir{{"results.csv", nil}}},
	}}
	try("create order", func() string {
		got := fmt.Sprint(CreateOrder(thesis))
		want := "[thesis thesis/chapters thesis/chapters/intro.tex thesis/chapters/figures thesis/chapters/figures/plot.svg thesis/data thesis/data/results.csv]"
		if got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	})
	try("delete order", func() string {
		got := fmt.Sprint(DeleteOrder(thesis))
		want := "[thesis/chapters/intro.tex thesis/chapters/figures/plot.svg thesis/chapters/figures thesis/chapters thesis/data/results.csv thesis/data thesis]"
		if got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	})
	try("a single file", func() string {
		f := &Dir{"notes.txt", nil}
		if fmt.Sprint(CreateOrder(f)) != "[notes.txt]" || fmt.Sprint(DeleteOrder(f)) != "[notes.txt]" {
			return fmt.Sprintf("CreateOrder = %v, DeleteOrder = %v for a single file", CreateOrder(f), DeleteOrder(f))
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
