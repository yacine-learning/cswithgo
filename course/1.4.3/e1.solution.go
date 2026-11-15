package main

import (
	"fmt"
	"sort"
	"strings"
)

// Plan groups services into deployment waves.
//
// deps[s] lists the services s depends on: each of them must be in an earlier wave than s.
// A service that appears only as a dependency is a service too.
//
//   - wave 0 holds every service with no dependency; wave k holds every service whose
//     dependencies are all in waves before k (each service as early as possible);
//   - the services of each wave are sorted by name;
//   - when some services can never be deployed (they are on a cycle, or depend on one),
//     Plan returns the error "cannot order: a, b, c": all those services, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	waiting := map[string]int{}       // number of dependencies not deployed yet
	neededBy := map[string][]string{} // the reverse graph
	for s, ds := range deps {
		waiting[s] += len(ds)
		for _, d := range ds {
			if _, ok := waiting[d]; !ok {
				waiting[d] = 0
			}
			neededBy[d] = append(neededBy[d], s)
		}
	}
	var wave []string
	for s, n := range waiting {
		if n == 0 {
			wave = append(wave, s)
		}
	}
	var waves [][]string
	placed := 0
	for len(wave) > 0 {
		sort.Strings(wave)
		waves = append(waves, wave)
		placed += len(wave)
		var next []string
		for _, s := range wave {
			for _, t := range neededBy[s] {
				waiting[t]--
				if waiting[t] == 0 {
					next = append(next, t)
				}
			}
		}
		wave = next
	}
	if placed < len(waiting) {
		var stuck []string
		for s, n := range waiting {
			if n > 0 {
				stuck = append(stuck, s)
			}
		}
		sort.Strings(stuck)
		return nil, fmt.Errorf("cannot order: %s", strings.Join(stuck, ", "))
	}
	return waves, nil
}
