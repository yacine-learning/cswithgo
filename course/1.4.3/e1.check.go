package main

import "fmt"

func expect(deps map[string][]string, want string) func() string {
	return func() string {
		for i := 0; i < 10; i++ { // several runs: map order must not leak into the result
			rounds, err := Plan(deps)
			if err != nil {
				return "unexpected error: " + err.Error()
			}
			if got := fmt.Sprint(rounds); got != want {
				return fmt.Sprintf("rounds %s, expected %s", got, want)
			}
		}
		return ""
	}
}

func expectErr(deps map[string][]string, want string) func() string {
	return func() string {
		rounds, err := Plan(deps)
		if err == nil {
			return fmt.Sprintf("no error (rounds %v); expected %q", rounds, want)
		}
		if err.Error() != want {
			return fmt.Sprintf("error %q, expected %q", err.Error(), want)
		}
		return ""
	}
}

func main() {
	try("a program in three rounds", expect(map[string][]string{
		"main": {"parser"}, "parser": {"bufio", "strings", "unicode"}, "eval": {"bufio", "unicode"},
	}, "[[bufio strings unicode] [eval parser] [main]]"))
	try("each package as early as possible", expect(map[string][]string{
		"a": {"b"}, "b": {"c"}, "x": {"c"},
	}, "[[c] [b x] [a]]"))
	try("packages that import nothing", expect(map[string][]string{"b": nil, "a": {}}, "[[a b]]"))
	try("an import shared by several packages", expect(map[string][]string{
		"a": {"unicode"}, "b": {"unicode"}, "c": {"a", "b", "unicode"},
	}, "[[unicode] [a b] [c]]"))
	try("nothing to build", expect(map[string][]string{}, "[]"))
	try("a cycle", expectErr(map[string][]string{"a": {"b"}, "b": {"c"}, "c": {"a"}}, "cannot order: a, b, c"))
	try("a package waiting behind a cycle is reported too", expectErr(map[string][]string{
		"main": {"parser"}, "parser": {"ast"}, "ast": {"types"}, "types": {"parser"}, "strings": {"unicode"},
	}, "cannot order: ast, main, parser, types"))
	try("a package that imports itself", expectErr(map[string][]string{"a": {"a"}, "b": nil}, "cannot order: a"))
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
