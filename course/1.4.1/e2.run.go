package main

import "fmt"

func main() {
	g := Build([]string{"main -> fmt", "fmt -> io", "parser->lexer", "main -> lexer"})
	for _, n := range []string{"main", "fmt", "parser", "lexer", "io"} {
		outs, ok := g[n]
		fmt.Printf("%-7s in graph: %-5v -> %v\n", n, ok, outs)
	}
}
