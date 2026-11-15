package main

import "fmt"

func same(got, want map[string]string) string {
	if fmt.Sprint(got) != fmt.Sprint(want) {
		return fmt.Sprintf("got %v\n      expected %v", got, want)
	}
	return ""
}

func main() {
	try("a flat document stays flat", func() string {
		return same(Flatten(map[string]any{"name": "main", "lines": 1}), map[string]string{"name": "main", "lines": "1"})
	})
	try("one level of nesting", func() string {
		return same(Flatten(map[string]any{"pos": map[string]any{"line": 3, "col": 7}}),
			map[string]string{"pos.line": "3", "pos.col": "7"})
	})
	try("deep nesting", func() string {
		return same(Flatten(map[string]any{"a": map[string]any{"b": map[string]any{"c": map[string]any{"d": true}}}}),
			map[string]string{"a.b.c.d": "true"})
	})
	try("lists use indexes", func() string {
		return same(Flatten(map[string]any{"imports": []any{"fmt", "os"}}), map[string]string{"imports.0": "fmt", "imports.1": "os"})
	})
	try("maps inside lists", func() string {
		return same(Flatten(map[string]any{"funcs": []any{map[string]any{"name": "main"}, map[string]any{"name": "parse"}}}),
			map[string]string{"funcs.0.name": "main", "funcs.1.name": "parse"})
	})
	try("an empty document", func() string { return same(Flatten(map[string]any{}), map[string]string{}) })
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
