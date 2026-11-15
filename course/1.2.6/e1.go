package main

// Report summarises a trace.
type Report struct {
	Durations map[string]int // total time of each span name (a name can appear several times)
	MaxDepth  int            // deepest nesting: 1 when spans never nest
	Slowest   []string       // up to 3 names, longest total first; ties in name order
}

// Analyze reads a trace: lines "TIME start NAME" and "TIME end NAME", where TIME is an
// integer and never decreases. Spans nest: an "end" must close the innermost open span.
//
// It returns the first problem as an error that names the line (counting from 1):
//
//	"line 3: cannot read \"oops\""         a line that is not TIME start|end NAME
//	"line 5: end test but nothing is open"
//	"line 5: end test but build is open"   (build is the innermost open span)
//	"end of log: build is still open"      (the innermost span left open)
func Analyze(lines []string) (Report, error) {
	return Report{}, nil
}
