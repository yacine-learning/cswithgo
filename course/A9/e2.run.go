package main

import "fmt"

func main() {
	compiler := &Dir{"compiler", []*Dir{
		{"parser", []*Dir{{"parse.go", nil}, {"testdata", []*Dir{{"expr.txt", nil}}}}},
		{"lexer", []*Dir{{"lex.go", nil}}},
	}}
	fmt.Println("create (a directory before what it contains):")
	for _, p := range CreateOrder(compiler) {
		fmt.Println("  ", p)
	}
	fmt.Println("remove (what a directory contains before the directory):")
	for _, p := range DeleteOrder(compiler) {
		fmt.Println("  ", p)
	}
}
