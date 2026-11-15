package main

import "fmt"

func main() {
	var b Browser
	step := func(what string, ok bool) {
		fmt.Printf("%-14s ok=%-5v current=%-10s back=%v forward=%v\n", what, ok, b.Current, b.back.items, b.forward.items)
	}
	b.Visit("home")
	step("Visit home", true)
	b.Visit("docs")
	step("Visit docs", true)
	b.Visit("api")
	step("Visit api", true)
	step("Back", b.Back())
	step("Back", b.Back())
	step("Forward", b.Forward())
	b.Visit("blog")
	step("Visit blog", true)
	step("Forward", b.Forward())
}
