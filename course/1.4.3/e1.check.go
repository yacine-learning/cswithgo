package main

import "fmt"

func expect(deps map[string][]string, want string) func() string {
	return func() string {
		for i := 0; i < 10; i++ { // several runs: map order must not leak into the result
			waves, err := Plan(deps)
			if err != nil {
				return "unexpected error: " + err.Error()
			}
			if got := fmt.Sprint(waves); got != want {
				return fmt.Sprintf("waves %s, expected %s", got, want)
			}
		}
		return ""
	}
}

func expectErr(deps map[string][]string, want string) func() string {
	return func() string {
		waves, err := Plan(deps)
		if err == nil {
			return fmt.Sprintf("no error (waves %v); expected %q", waves, want)
		}
		if err.Error() != want {
			return fmt.Sprintf("error %q, expected %q", err.Error(), want)
		}
		return ""
	}
}

func main() {
	try("a release in three waves", expect(map[string][]string{
		"web": {"api"}, "api": {"db", "cache", "queue"}, "worker": {"db", "queue"},
	}, "[[cache db queue] [api worker] [web]]"))
	try("each service as early as possible", expect(map[string][]string{
		"a": {"b"}, "b": {"c"}, "x": {"c"},
	}, "[[c] [b x] [a]]"))
	try("services with no dependencies at all", expect(map[string][]string{"b": nil, "a": {}}, "[[a b]]"))
	try("a dependency listed twice by different services", expect(map[string][]string{
		"a": {"db"}, "b": {"db"}, "c": {"a", "b", "db"},
	}, "[[db] [a b] [c]]"))
	try("nothing to deploy", expect(map[string][]string{}, "[]"))
	try("a cycle", expectErr(map[string][]string{"a": {"b"}, "b": {"c"}, "c": {"a"}}, "cannot order: a, b, c"))
	try("a service waiting behind a cycle is reported too", expectErr(map[string][]string{
		"web": {"api"}, "api": {"auth"}, "auth": {"users"}, "users": {"api"}, "cron": {"db"},
	}, "cannot order: api, auth, users, web"))
	try("a service that depends on itself", expectErr(map[string][]string{"a": {"a"}, "b": nil}, "cannot order: a"))
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
