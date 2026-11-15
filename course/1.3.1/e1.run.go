package main

import "fmt"

func main() {
	thesis := &Dir{Name: "thesis", Files: []File{{"abstract.txt", 2}}, Subdirs: []*Dir{
		{Name: "chapters", Files: []File{{"intro.tex", 12}, {"methods.tex", 5}}, Subdirs: []*Dir{
			{Name: "figures", Files: []File{{"plot.svg", 20}}},
		}},
		{Name: "data", Files: []File{{"results.csv", 1}}},
	}}
	fmt.Println("\ntotal:", DirSize(thesis, 0))
}
