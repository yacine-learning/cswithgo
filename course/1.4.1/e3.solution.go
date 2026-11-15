package main

import "sort"

// Graph maps each species to the species it eats.
type Graph map[string][]string

// Producers returns the species that eat nothing, sorted.
func Producers(g Graph) []string {
	var out []string
	for s, prey := range g {
		if len(prey) == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}

// ApexPredators returns the species that no other species eats, sorted.
func ApexPredators(g Graph) []string {
	eaten := map[string]int{} // in-degree: how many species eat it
	for _, prey := range g {
		for _, p := range prey {
			eaten[p]++
		}
	}
	var out []string
	for s := range g {
		if eaten[s] == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}
