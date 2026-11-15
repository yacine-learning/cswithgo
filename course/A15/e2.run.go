package main

import (
	"fmt"
	"strings"
)

func main() {
	keywords := strings.Fields(`break case chan const continue default defer else fallthrough for func go goto
	if import interface map package range return select struct switch type var`)
	for _, typo := range []string{"fucn", "retrun", "pakage", "swich", "contine", "defr", "esle", "for", "xyz", "imprt"} {
		if s, ok := Suggest(typo, keywords); ok {
			fmt.Printf("%-8s did you mean %q? (distance %d)\n", typo, s, EditDistance(typo, s))
		} else {
			fmt.Printf("%-8s no suggestion\n", typo)
		}
	}
}
