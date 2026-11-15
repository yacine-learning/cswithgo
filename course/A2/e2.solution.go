package main

import "fmt"

// Flatten turns a nested document into flat paths joined by dots:
// {"pos": {"line": 3, "col": 7}} -> {"pos.col": "7", "pos.line": "3"}.
// Lists use the index of each element as its key: {"imports": ["fmt", "os"]} -> {"imports.0": "fmt", "imports.1": "os"}.
// Leaf values are written with fmt.Sprint.
func Flatten(doc map[string]any) map[string]string {
	out := map[string]string{}
	for k, v := range doc {
		walk(k, v, out)
	}
	return out
}

func walk(prefix string, v any, out map[string]string) {
	switch x := v.(type) {
	case map[string]any:
		for k, child := range x {
			walk(prefix+"."+k, child, out)
		}
	case []any:
		for i, child := range x {
			walk(fmt.Sprintf("%s.%d", prefix, i), child, out)
		}
	default:
		out[prefix] = fmt.Sprint(v)
	}
}
