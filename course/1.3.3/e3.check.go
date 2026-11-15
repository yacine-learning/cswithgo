package main

import "fmt"

func main() {
	app := &Dir{"app", []*Dir{
		{"src", []*Dir{{"main.go", nil}, {"api", []*Dir{{"api.go", nil}}}}},
		{"config", []*Dir{{"app.yml", nil}}},
	}}
	try("create order", func() string {
		got := fmt.Sprint(CreateOrder(app))
		want := "[app app/src app/src/main.go app/src/api app/src/api/api.go app/config app/config/app.yml]"
		if got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	})
	try("delete order", func() string {
		got := fmt.Sprint(DeleteOrder(app))
		want := "[app/src/main.go app/src/api/api.go app/src/api app/src app/config/app.yml app/config app]"
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
