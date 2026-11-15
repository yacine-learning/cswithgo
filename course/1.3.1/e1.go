package main

import (
	"fmt"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name    string
	Files   []File
	Subdirs []*Dir
}

// DirSize returns the size of d: its files plus each subdirectory.
// depth is only used to indent the trace.
func DirSize(d *Dir, depth int) int {
	pad := strings.Repeat("  ", depth)
	fmt.Printf("%s> enter %s\n", pad, d.Name)
	total := 0
	for _, f := range d.Files {
		total += f.Size
	}
	for _, sub := range d.Subdirs {
		fmt.Printf("%s  %s pauses: call DirSize(%s)\n", pad, d.Name, sub.Name)
		total += DirSize(sub, depth+1) // this call pauses here until the inner one returns
		fmt.Printf("%s  %s resumes, total so far %d\n", pad, d.Name, total)
	}
	fmt.Printf("%s< leave %s = %d\n", pad, d.Name, total)
	return total
}
