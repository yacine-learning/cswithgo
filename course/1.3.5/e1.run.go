package main

import "fmt"

func main() {
	files := []string{
		"thesis/abstract.txt 2",
		"thesis/chapters/intro.tex 12",
		"thesis/chapters/figures/plot.svg 20",
		"thesis/data/results.csv 1",
		"notes/2026-11-14.txt 300",
		"notes/archive/2026-10.txt 900",
	}
	for _, d := range []int{0, 1, 2} {
		out, err := Sizes(files, d)
		fmt.Printf("depth %d:\n", d)
		if err != nil {
			fmt.Println("error:", err)
		}
		for _, l := range out {
			fmt.Println(l)
		}
		fmt.Println()
	}
}
