package main

import "fmt"

func main() {
	base := make([]string, 0, 10) // room for more: as when a slice was built with append
	base = append(base, "deploy", "--env=prod")
	canary := WithFlag(base, "--canary")
	full := WithFlag(base, "--all-regions")
	fmt.Println("canary:", canary)
	fmt.Println("full:  ", full)
	fmt.Println("base:  ", base)
}
