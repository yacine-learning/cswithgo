package main

import "fmt"

func expect(deps map[string][]string, want string) func() string {
	return func() string {
		for i := 0; i < 10; i++ { // several runs: map order must not leak into the result
			semesters, err := Plan(deps)
			if err != nil {
				return "unexpected error: " + err.Error()
			}
			if got := fmt.Sprint(semesters); got != want {
				return fmt.Sprintf("semesters %s, expected %s", got, want)
			}
		}
		return ""
	}
}

func expectErr(deps map[string][]string, want string) func() string {
	return func() string {
		semesters, err := Plan(deps)
		if err == nil {
			return fmt.Sprintf("no error (semesters %v); expected %q", semesters, want)
		}
		if err.Error() != want {
			return fmt.Sprintf("error %q, expected %q", err.Error(), want)
		}
		return ""
	}
}

func main() {
	try("a curriculum in three semesters", expect(map[string][]string{
		"physics": {"calculus"}, "calculus": {"arithmetic", "geometry", "logic"}, "probability": {"arithmetic", "logic"},
	}, "[[arithmetic geometry logic] [calculus probability] [physics]]"))
	try("each topic as early as possible", expect(map[string][]string{
		"a": {"b"}, "b": {"c"}, "x": {"c"},
	}, "[[c] [b x] [a]]"))
	try("topics with no prerequisites at all", expect(map[string][]string{"b": nil, "a": {}}, "[[a b]]"))
	try("a prerequisite shared by several topics", expect(map[string][]string{
		"a": {"logic"}, "b": {"logic"}, "c": {"a", "b", "logic"},
	}, "[[logic] [a b] [c]]"))
	try("nothing to study", expect(map[string][]string{}, "[]"))
	try("a cycle", expectErr(map[string][]string{"a": {"b"}, "b": {"c"}, "c": {"a"}}, "cannot order: a, b, c"))
	try("a topic waiting behind a cycle is reported too", expectErr(map[string][]string{
		"physics": {"calculus"}, "calculus": {"limits"}, "limits": {"series"}, "series": {"calculus"}, "geometry": {"logic"},
	}, "cannot order: calculus, limits, physics, series"))
	try("a topic that needs itself", expectErr(map[string][]string{"a": {"a"}, "b": nil}, "cannot order: a"))
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
