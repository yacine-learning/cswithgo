package main

import "fmt"

func main() {
	calc := Graph{
		"main":    {"parser", "eval"},
		"parser":  {"lexer", "ast"},
		"eval":    {"lexer", "errors"},
		"lexer":   {"token"},
		"token":   {"errors"},
		"ast":     {"errors", "unicode"},
		"errors":  nil,
		"unicode": nil,
	}
	fmt.Println("leaves:")
	fmt.Println("  ", Leaves(calc))
	fmt.Println("roots:")
	fmt.Println("  ", Roots(calc))
}
