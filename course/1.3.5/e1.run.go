package main

import "fmt"

func main() {
	files := []string{
		"compiler/go.mod 2",
		"compiler/lexer/lex.go 12",
		"compiler/lexer/testdata/input.txt 20",
		"compiler/parser/parse.go 1",
		"notes/2026-11-14.txt 300",
		"notes/archive/2026-10.txt 900",
	}
	for _, d := range []int{0, 1, 2} {
		out, err := Sizes(files, d)
		fmt.Printf("depth %d:\n", d)
		if err != nil {
			fmt.Println("error:", err)
		}
		for _, l := range out {
			fmt.Println(l)
		}
		fmt.Println()
	}
}
