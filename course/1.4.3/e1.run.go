package main

import "fmt"

func show(name string, imports map[string][]string) {
	fmt.Println(name + ":")
	rounds, err := Plan(imports)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	for i, r := range rounds {
		fmt.Printf("  round %d: %v\n", i, r)
	}
}

func main() {
	show("a small calculator", map[string][]string{
		"main":   {"parser", "eval"},
		"parser": {"lexer", "ast"},
		"eval":   {"ast", "strconv"},
		"lexer":  {"token", "unicode"},
		"ast":    {"token"},
	})
	show("with a cycle", map[string][]string{
		"main":   {"parser"},
		"parser": {"ast"},
		"ast":    {"types"},
		"types":  {"parser"},
		"fmt":    {"strconv"},
	})
}
