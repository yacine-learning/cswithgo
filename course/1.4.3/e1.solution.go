package main

import (
	"fmt"
	"sort"
	"strings"
)

// Plan groups packages into build rounds.
//
// deps[p] lists the packages that p imports: each of them must be built in an earlier
// round than p. A package that appears only as an import is a package too.
//
//   - round 0 holds every package that imports nothing; round k holds every package whose
//     imports are all built in rounds before k (each package as early as possible);
//   - the packages of each round are sorted by name;
//   - when some packages can never be built (they are on a cycle, or import one),
//     Plan returns the error "cannot order: a, b, c": all those packages, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	waiting := map[string]int{}       // number of imports not built yet
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
	var round []string
	for s, n := range waiting {
		if n == 0 {
			round = append(round, s)
		}
	}
	var rounds [][]string
	placed := 0
	for len(round) > 0 {
		sort.Strings(round)
		rounds = append(rounds, round)
		placed += len(round)
		var next []string
		for _, s := range round {
			for _, t := range neededBy[s] {
				waiting[t]--
				if waiting[t] == 0 {
					next = append(next, t)
				}
			}
		}
		round = next
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
	return rounds, nil
}
