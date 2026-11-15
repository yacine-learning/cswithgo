package main

import "sort"

// Graph maps each service to the services it depends on.
type Graph map[string][]string

// DeployFirst returns the services that depend on nothing, sorted.
func DeployFirst(g Graph) []string {
	var out []string
	for s, deps := range g {
		if len(deps) == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}

// SafeToRemove returns the services that no other service depends on, sorted.
func SafeToRemove(g Graph) []string {
	needed := map[string]int{} // in-degree: how many services depend on it
	for _, deps := range g {
		for _, d := range deps {
			needed[d]++
		}
	}
	var out []string
	for s := range g {
		if needed[s] == 0 {
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}
