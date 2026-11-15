package main

import "fmt"

func main() {
	compiler := &Dir{"compiler", []*Dir{
		{"parser", []*Dir{{"parse.go", nil}, {"testdata", []*Dir{{"expr.txt", nil}}}}},
		{"lexer", []*Dir{{"lex.go", nil}}},
	}}
	try("create order", func() string {
		got := fmt.Sprint(CreateOrder(compiler))
		want := "[compiler compiler/parser compiler/parser/parse.go compiler/parser/testdata compiler/parser/testdata/expr.txt compiler/lexer compiler/lexer/lex.go]"
		if got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	})
	try("delete order", func() string {
		got := fmt.Sprint(DeleteOrder(compiler))
		want := "[compiler/parser/parse.go compiler/parser/testdata/expr.txt compiler/parser/testdata compiler/parser compiler/lexer/lex.go compiler/lexer compiler]"
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
