package main

// Flatten turns a nested document into flat paths joined by dots:
// {"pos": {"line": 3, "col": 7}} -> {"pos.col": "7", "pos.line": "3"}.
// Lists use the index of each element as its key: {"imports": ["fmt", "os"]} -> {"imports.0": "fmt", "imports.1": "os"}.
// Leaf values are written with fmt.Sprint.
func Flatten(doc map[string]any) map[string]string {
	out := map[string]string{}
	return out
}
