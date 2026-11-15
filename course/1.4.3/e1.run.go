package main

import "fmt"

func show(name string, deps map[string][]string) {
	fmt.Println(name + ":")
	semesters, err := Plan(deps)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	for i, s := range semesters {
		fmt.Printf("  semester %d: %v\n", i, s)
	}
}

func main() {
	show("a mathematics curriculum", map[string][]string{
		"calculus":    {"functions", "algebra"},
		"functions":   {"algebra"},
		"probability": {"algebra", "combinatorics"},
		"statistics":  {"probability", "calculus"},
	})
	show("with a cycle", map[string][]string{
		"physics":  {"calculus"},
		"calculus": {"limits"},
		"limits":   {"series"},
		"series":   {"calculus"},
		"geometry": {"logic"},
	})
}
