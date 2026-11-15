package main

import "fmt"

func main() {
	app := &Dir{"app", []*Dir{
		{"src", []*Dir{{"main.go", nil}, {"api", []*Dir{{"api.go", nil}}}}},
		{"config", []*Dir{{"app.yml", nil}}},
	}}
	fmt.Println("create (mkdir, then the files inside):")
	for _, p := range CreateOrder(app) {
		fmt.Println("  ", p)
	}
	fmt.Println("delete (like rm -rf does it):")
	for _, p := range DeleteOrder(app) {
		fmt.Println("  ", p)
	}
}
