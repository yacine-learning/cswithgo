package main

import "fmt"

func main() {
	for _, s := range []string{
		`{"ports": [80, 443], "env": {"tier": "prod"}}`,
		`{"ports": [80, 443}, "env": {}}`,
		`resource "x" { tags = { a = "b" }`,
		`list = [1, 2])`,
	} {
		ok, pos := Balanced(s)
		fmt.Printf("%v %3d  %s\n", ok, pos, s)
		if !ok && pos >= 0 && pos < len(s) {
			fmt.Printf("         %*s^\n", pos+1, "")
		}
	}
}
