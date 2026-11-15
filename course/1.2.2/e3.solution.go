package main

import (
	"sort"
	"strings"
)

// Count is how many log lines had one status code.
type Count struct {
	Code string
	N    int
}

// StatusCounts counts the status codes of access-log lines such as
// "GET /api/users 200 12ms", most frequent first, then by code.
func StatusCounts(lines []string) []Count {
	counts := map[string]int{}
	for _, l := range lines {
		f := strings.Fields(l)
		if len(f) >= 3 {
			counts[f[2]]++
		}
	}
	out := make([]Count, 0, len(counts))
	for code, n := range counts {
		out = append(out, Count{code, n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].N != out[j].N {
			return out[i].N > out[j].N
		}
		return out[i].Code < out[j].Code
	})
	return out
}
