package main

// ExternalSort sorts lines as if only memory lines (memory >= 1) fit in RAM at once:
//
//  1. cut lines into runs of memory lines (the last one may be shorter) and sort each run;
//  2. merge the runs two by two (the 1st with the 2nd, the 3rd with the 4th, …; an odd
//     last run goes on unchanged to the next round), until one run is left.
//
// It returns the sorted lines, the number of runs made in step 1, and the number of
// merge rounds in step 2. The caller's slice is not changed. No lines: nil, 0, 0.
func ExternalSort(lines []string, memory int) (sorted []string, runs, rounds int) {
	return nil, 0, 0
}
