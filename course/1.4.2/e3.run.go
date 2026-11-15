package main

import "fmt"

func main() {
	ok := Graph{"main": {"fmt", "parser"}, "parser": {"lexer", "token"}, "fmt": nil, "lexer": nil, "token": nil}
	cyclic := Graph{"parser": {"ast"}, "ast": {"types"}, "types": {"parser", "errors"}, "errors": nil, "main": {"parser"}}
	fmt.Println("no cycle:  ", FindCycle(ok))
	fmt.Println("with cycle:", FindCycle(cyclic))
}
