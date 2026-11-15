package main

import "fmt"

func main() {
	// elements in order of atomic number, keyed by their period (row of the periodic table)
	elements := []Rec{{1, "H"}, {1, "He"}, {2, "Li"}, {2, "Be"}, {3, "Na"}, {3, "Mg"}, {2, "B"}, {3, "Al"}}
	fmt.Println("by period:", MergeSort(elements))
}
