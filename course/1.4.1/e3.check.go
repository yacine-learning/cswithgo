package main

import "fmt"

func main() {
	g := Graph{
		"main": {"parser", "eval"}, "parser": {"lexer", "ast"}, "eval": {"lexer", "errors"},
		"repl": {"eval"}, "lexer": {"token"}, "token": {"errors"}, "ast": {"errors", "unicode"},
		"errors": nil, "unicode": nil,
	}
	try("leaves (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(Leaves(g)); got != "[errors unicode]" {
				return "Leaves = " + got + ", expected [errors unicode]"
			}
		}
		return ""
	})
	try("roots (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(Roots(g)); got != "[main repl]" {
				return "Roots = " + got + ", expected [main repl]"
			}
		}
		return ""
	})
	try("a lone package is both", func() string {
		g := Graph{"hello": nil}
		if fmt.Sprint(Leaves(g)) != "[hello]" || fmt.Sprint(Roots(g)) != "[hello]" {
			return fmt.Sprintf("Leaves = %v, Roots = %v, expected [hello] for both", Leaves(g), Roots(g))
		}
		return ""
	})
	try("a chain main -> parser -> lexer -> unicode", func() string {
		g := Graph{"main": {"parser"}, "parser": {"lexer"}, "lexer": {"unicode"}, "unicode": nil}
		if fmt.Sprint(Leaves(g)) != "[unicode]" || fmt.Sprint(Roots(g)) != "[main]" {
			return fmt.Sprintf("Leaves = %v, Roots = %v; expected [unicode] and [main]", Leaves(g), Roots(g))
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
