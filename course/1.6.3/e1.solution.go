package main

import "slices"

// ExternalSort sorts lines as if only memory lines (memory >= 1) fit in RAM at once:
//
//  1. cut lines into runs of memory lines (the last one may be shorter) and sort each run;
//  2. merge the runs two by two (the 1st with the 2nd, the 3rd with the 4th, …; an odd
//     last run goes on unchanged to the next round), until one run is left.
//
// It returns the sorted lines, the number of runs made in step 1, and the number of
// merge rounds in step 2. The caller's slice is not changed. No lines: nil, 0, 0.
func ExternalSort(lines []string, memory int) (sorted []string, runs, rounds int) {
	var rs [][]string
	for i := 0; i < len(lines); i += memory {
		run := slices.Clone(lines[i:min(i+memory, len(lines))])
		slices.Sort(run)
		rs = append(rs, run)
	}
	runs = len(rs)
	for len(rs) > 1 {
		var next [][]string
		for i := 0; i < len(rs); i += 2 {
			if i+1 < len(rs) {
				next = append(next, merge(rs[i], rs[i+1]))
			} else {
				next = append(next, rs[i])
			}
		}
		rs = next
		rounds++
	}
	if runs == 0 {
		return nil, 0, 0
	}
	return rs[0], runs, rounds
}

func merge(a, b []string) []string {
	out := make([]string, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	return append(out, b[j:]...)
}
