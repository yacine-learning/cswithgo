package main

import "fmt"

func main() {
	// processes in order of arrival, keyed by their priority (1 runs first)
	procs := []Rec{{2, "cc"}, {1, "edit"}, {3, "troff"}, {2, "ls"}, {1, "sh"}, {3, "spell"}, {2, "mail"}, {3, "lpr"}}
	fmt.Println("by priority, first come first served within each:", MergeSort(procs))
}
