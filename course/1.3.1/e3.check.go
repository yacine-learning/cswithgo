package main

import "fmt"

func same(got, want map[string]string) string {
	if fmt.Sprint(got) != fmt.Sprint(want) {
		return fmt.Sprintf("got %v\n      expected %v", got, want)
	}
	return ""
}

func main() {
	try("a flat config stays flat", func() string {
		return same(Flatten(map[string]any{"a": 1, "b": "x"}), map[string]string{"a": "1", "b": "x"})
	})
	try("one level of nesting", func() string {
		return same(Flatten(map[string]any{"db": map[string]any{"host": "x", "port": 5432}}),
			map[string]string{"db.host": "x", "db.port": "5432"})
	})
	try("deep nesting", func() string {
		return same(Flatten(map[string]any{"a": map[string]any{"b": map[string]any{"c": map[string]any{"d": true}}}}),
			map[string]string{"a.b.c.d": "true"})
	})
	try("lists use indexes", func() string {
		return same(Flatten(map[string]any{"ports": []any{80, 443}}), map[string]string{"ports.0": "80", "ports.1": "443"})
	})
	try("maps inside lists", func() string {
		return same(Flatten(map[string]any{"users": []any{map[string]any{"name": "ana"}, map[string]any{"name": "bo"}}}),
			map[string]string{"users.0.name": "ana", "users.1.name": "bo"})
	})
	try("an empty config", func() string { return same(Flatten(map[string]any{}), map[string]string{}) })
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
