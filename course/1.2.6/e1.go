package main

// Report summarises a call trace.
type Report struct {
	Durations map[string]int // total time of each function (a function can be called several times)
	MaxDepth  int            // deepest nesting: 1 when calls never nest
	Slowest   []string       // up to 3 names, longest total first; ties in name order
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
	return Report{}, nil
}
