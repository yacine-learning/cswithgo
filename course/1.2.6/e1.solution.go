package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Report summarises a call trace.
type Report struct {
	Durations map[string]int // total time of each function (a function can be called several times)
	MaxDepth  int            // deepest nesting: 1 when calls never nest
	Slowest   []string       // up to 3 names, longest total first; ties in name order
}

type open struct {
	name  string
	start int
}

// Analyze reads a trace: lines "TIME start NAME" and "TIME end NAME", where TIME is an
// integer and never decreases. Calls nest: an "end" must close the innermost open call.
//
// It returns the first problem as an error that names the line (counting from 1):
//
//	"line 3: cannot read \"oops\""         a line that is not TIME start|end NAME
//	"line 5: end lex but nothing is open"
//	"line 5: end lex but parse is open"   (parse is the innermost open call)
//	"end of trace: parse is still open"      (the innermost call left open)
func Analyze(lines []string) (Report, error) {
	r := Report{Durations: map[string]int{}}
	var stack []open
	for i, line := range lines {
		f := strings.Fields(line)
		t, err := 0, error(nil)
		if len(f) == 3 {
			t, err = strconv.Atoi(f[0])
		}
		if len(f) != 3 || err != nil || (f[1] != "start" && f[1] != "end") {
			return Report{}, fmt.Errorf("line %d: cannot read %q", i+1, line)
		}
		name := f[2]
		if f[1] == "start" {
			stack = append(stack, open{name, t})
			r.MaxDepth = max(r.MaxDepth, len(stack))
			continue
		}
		if len(stack) == 0 {
			return Report{}, fmt.Errorf("line %d: end %s but nothing is open", i+1, name)
		}
		top := stack[len(stack)-1]
		if top.name != name {
			return Report{}, fmt.Errorf("line %d: end %s but %s is open", i+1, name, top.name)
		}
		stack = stack[:len(stack)-1]
		r.Durations[name] += t - top.start
	}
	if len(stack) > 0 {
		return Report{}, fmt.Errorf("end of trace: %s is still open", stack[len(stack)-1].name)
	}
	for name := range r.Durations {
		r.Slowest = append(r.Slowest, name)
	}
	sort.Slice(r.Slowest, func(a, b int) bool {
		da, db := r.Durations[r.Slowest[a]], r.Durations[r.Slowest[b]]
		if da != db {
			return da > db
		}
		return r.Slowest[a] < r.Slowest[b]
	})
	if len(r.Slowest) > 3 {
		r.Slowest = r.Slowest[:3]
	}
	return r, nil
}
