package main

import (
	"fmt"
	"sort"
	"strings"
)

// Plan groups topics into semesters.
//
// deps[t] lists the prerequisites of topic t: each of them must be in an earlier semester
// than t. A topic that appears only as a prerequisite is a topic too.
//
//   - semester 0 holds every topic with no prerequisite; semester k holds every topic whose
//     prerequisites are all in semesters before k (each topic as early as possible);
//   - the topics of each semester are sorted by name;
//   - when some topics can never be studied (they are on a cycle, or need one),
//     Plan returns the error "cannot order: a, b, c": all those topics, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	waiting := map[string]int{}       // number of prerequisites not studied yet
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
	var semester []string
	for s, n := range waiting {
		if n == 0 {
			semester = append(semester, s)
		}
	}
	var semesters [][]string
	placed := 0
	for len(semester) > 0 {
		sort.Strings(semester)
		semesters = append(semesters, semester)
		placed += len(semester)
		var next []string
		for _, s := range semester {
			for _, t := range neededBy[s] {
				waiting[t]--
				if waiting[t] == 0 {
					next = append(next, t)
				}
			}
		}
		semester = next
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
	return semesters, nil
}
