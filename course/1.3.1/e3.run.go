package main

import (
	"fmt"
	"sort"
)

func main() {
	config := map[string]any{
		"service": "api",
		"db":      map[string]any{"host": "db.internal", "port": 5432, "tls": map[string]any{"enabled": true}},
		"ports":   []any{80, 443},
	}
	flat := Flatten(config)
	keys := make([]string, 0, len(flat))
	for k := range flat {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%s\n", k, flat[k])
	}
}
