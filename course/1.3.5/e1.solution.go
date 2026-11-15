package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	files    int // sizes of the files directly in this folder
	children map[string]*dir
}

func newDir() *dir { return &dir{children: map[string]*dir{}} }

// Sizes reads lines "PATH SIZE" (one per file; PATH has no spaces, SIZE is an integer)
// and returns one line "SIZE\tPATH" per folder:
//
//   - a folder's SIZE is the total of all files below it, at any depth;
//   - a top-level folder (the first part of a path) has depth 0, its subfolders 1, …;
//     only folders with depth <= maxDepth are listed;
//   - order: post-order: inside a folder, its subfolders first, sorted by
//     name, then the folder itself; top-level folders sorted by name;
//   - a file at the top level (a PATH without "/") is in no folder and is not listed;
//   - a line that cannot be read gives the error "line N: cannot read \"...\"" (N from 1).
func Sizes(lines []string, maxDepth int) ([]string, error) {
	root := newDir() // holds the top-level folders; never printed
	for i, line := range lines {
		f := strings.Fields(line)
		if len(f) != 2 {
			return nil, fmt.Errorf("line %d: cannot read %q", i+1, line)
		}
		size, err := strconv.Atoi(f[1])
		if err != nil {
			return nil, fmt.Errorf("line %d: cannot read %q", i+1, line)
		}
		parts := strings.Split(f[0], "/")
		d := root
		for _, name := range parts[:len(parts)-1] {
			if d.children[name] == nil {
				d.children[name] = newDir()
			}
			d = d.children[name]
		}
		d.files += size
	}
	var out []string
	var walk func(d *dir, path string, depth int) int
	walk = func(d *dir, path string, depth int) int {
		total := d.files
		for _, name := range sortedNames(d) {
			total += walk(d.children[name], path+"/"+name, depth+1)
		}
		if depth <= maxDepth {
			out = append(out, fmt.Sprintf("%d\t%s", total, path))
		}
		return total
	}
	for _, name := range sortedNames(root) {
		walk(root.children[name], name, 0)
	}
	return out, nil
}

func sortedNames(d *dir) []string {
	names := make([]string, 0, len(d.children))
	for n := range d.children {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}
