package main

import (
	"fmt"
	"strings"
)

func lines(s string) []string { return strings.Split(strings.TrimSpace(s), "\n") }

func errorCase(trace, want string) func() string {
	return func() string {
		_, err := Analyze(lines(trace))
		if err == nil {
			return fmt.Sprintf("no error; expected %q", want)
		}
		if err.Error() != want {
			return fmt.Sprintf("error %q, expected %q", err.Error(), want)
		}
		return ""
	}
}

func main() {
	good := `
10 start main
12 start parse
15 start lex
41 end lex
44 end parse
45 start eval
60 end eval
61 end main`
	try("durations of a nested trace", func() string {
		r, err := Analyze(lines(good))
		if err != nil {
			return "unexpected error: " + err.Error()
		}
		want := map[string]int{"main": 51, "parse": 32, "lex": 26, "eval": 15}
		if fmt.Sprint(r.Durations) != fmt.Sprint(want) {
			return fmt.Sprintf("durations %v, expected %v", r.Durations, want)
		}
		return ""
	})
	try("deepest nesting", func() string {
		r, _ := Analyze(lines(good))
		if r.MaxDepth != 3 {
			return fmt.Sprintf("MaxDepth %d, expected 3 (main > parse > lex)", r.MaxDepth)
		}
		return ""
	})
	try("the three slowest", func() string {
		r, _ := Analyze(lines(good))
		if fmt.Sprint(r.Slowest) != "[main parse lex]" {
			return fmt.Sprintf("Slowest %v, expected [main parse lex]", r.Slowest)
		}
		return ""
	})
	try("a name that appears twice adds up; ties in name order", func() string {
		r, err := Analyze(lines(`
0 start sqrt
5 end sqrt
5 start log
9 end log
9 start sqrt
13 end sqrt
13 start sin
22 end sin`))
		if err != nil {
			return "unexpected error: " + err.Error()
		}
		if r.Durations["sqrt"] != 9 {
			return fmt.Sprintf("sqrt ran 5 + 4: Durations[sqrt] = %d, expected 9", r.Durations["sqrt"])
		}
		if r.MaxDepth != 1 {
			return fmt.Sprintf("spans that never nest: MaxDepth %d, expected 1", r.MaxDepth)
		}
		if fmt.Sprint(r.Slowest) != "[sin sqrt log]" {
			return fmt.Sprintf("Slowest %v, expected [sin sqrt log] (both 9: sin before sqrt)", r.Slowest)
		}
		return ""
	})
	try("an empty trace", func() string {
		r, err := Analyze(nil)
		if err != nil || len(r.Durations) != 0 || r.MaxDepth != 0 || len(r.Slowest) != 0 {
			return fmt.Sprintf("Analyze(nil) = %+v, %v", r, err)
		}
		return ""
	})
	try("end closes the wrong span", errorCase("1 start parse\n2 start lex\n3 end parse", "line 3: end parse but lex is open"))
	try("end with nothing open", errorCase("1 start a\n2 end a\n3 end a", "line 3: end a but nothing is open"))
	try("a span left open", errorCase("1 start parse\n2 start lex\n3 end lex", "end of trace: parse is still open"))
	try("the innermost open span is named", errorCase("1 start parse\n2 start lex", "end of trace: lex is still open"))
	try("a line that cannot be read", errorCase("1 start a\noops\n3 end a", "line 2: cannot read \"oops\""))
	try("a time that is not a number", errorCase("x start a", "line 1: cannot read \"x start a\""))
	try("an unknown word", errorCase("1 begin a", "line 1: cannot read \"1 begin a\""))
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
