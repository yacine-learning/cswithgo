package main

import (
	"fmt"
	"strings"
)

func main() {
	// Go's keywords and predeclared names: what an editor offers as you type
	names := strings.Fields(`break case chan const continue default defer else fallthrough for func go goto
	if import interface map package range return select struct switch type var
	append bool byte cap clear close complex copy delete error false float64 imag int len make max min
	new nil panic print println real recover rune string true uint`)
	t := NewTrie()
	for _, n := range names {
		t.Insert(n)
	}
	for _, p := range []string{"co", "c", "go", "re", "f", "x"} {
		fmt.Printf("%-3s → %v\n", p, Complete(t, p))
	}
}
