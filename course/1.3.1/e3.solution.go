package main

import "fmt"

// Flatten turns a nested config into flat keys joined by dots:
// {"db": {"host": "x", "port": 5432}} -> {"db.host": "x", "db.port": "5432"}.
// Lists use the index of each element as its key: {"ports": [80, 443]} -> {"ports.0": "80", "ports.1": "443"}.
// Leaf values are written with fmt.Sprint.
func Flatten(config map[string]any) map[string]string {
	out := map[string]string{}
	for k, v := range config {
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
