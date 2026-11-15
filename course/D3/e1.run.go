package main

import "fmt"

func main() {
	for _, s := range []string{
		`f(x) = [3 * (x + 1)] / {2 - x}`,
		`f(x) = [3 * (x + 1)} / 2`,
		`(define (square x) (* x x)`,
		`v = [1, 2])`,
	} {
		ok, pos := Balanced(s)
		fmt.Printf("%v %3d  %s\n", ok, pos, s)
		if !ok && pos >= 0 && pos < len(s) {
			fmt.Printf("         %*s^\n", pos+1, "")
		}
	}
}
