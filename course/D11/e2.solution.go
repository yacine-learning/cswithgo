package main

import "sort"

// Graph maps each package to the packages it imports.
type Graph map[string][]string

// Leaves returns the packages that import nothing, sorted.
func Leaves(g Graph) []string {
	var out []string
	for s, imports := range g {
		if len(imports) == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}

// Roots returns the packages that no other package imports, sorted.
func Roots(g Graph) []string {
	importers := map[string]int{} // in-degree: how many packages import it
	for _, imports := range g {
		for _, p := range imports {
			importers[p]++
		}
	}
	var out []string
	for s := range g {
		if importers[s] == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}
