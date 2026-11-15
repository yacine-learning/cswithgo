package main

import (
	"fmt"
	"sort"
)

func main() {
	// what a compiler front end might report about one Go file
	file := map[string]any{
		"file":    "main.go",
		"package": "main",
		"imports": []any{"fmt", "strings"},
		"funcs": []any{
			map[string]any{"name": "main", "pos": map[string]any{"line": 5, "col": 6}},
			map[string]any{"name": "count", "pos": map[string]any{"line": 12, "col": 6}, "params": []any{"text"}},
		},
	}
	flat := Flatten(file)
	keys := make([]string, 0, len(flat))
	for k := range flat {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%s\n", k, flat[k])
	}
}
