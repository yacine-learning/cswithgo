package main

import "fmt"

func main() {
	thesis := &Dir{"thesis", []*Dir{
		{"chapters", []*Dir{{"intro.tex", nil}, {"figures", []*Dir{{"plot.svg", nil}}}}},
		{"data", []*Dir{{"results.csv", nil}}},
	}}
	fmt.Println("create (a directory before what it contains):")
	for _, p := range CreateOrder(thesis) {
		fmt.Println("  ", p)
	}
	fmt.Println("remove (what a directory contains before the directory):")
	for _, p := range DeleteOrder(thesis) {
		fmt.Println("  ", p)
	}
}
