package main

import (
	"fmt"
	"sort"
)

func main() {
	jupiter := map[string]any{
		"name":  "Jupiter",
		"orbit": map[string]any{"period_days": 4333, "eccentricity": 0.0489, "around": map[string]any{"star": "Sun"}},
		"moons": []any{"Io", "Europa", "Ganymede", "Callisto"},
	}
	flat := Flatten(jupiter)
	keys := make([]string, 0, len(flat))
	for k := range flat {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%s\n", k, flat[k])
	}
}
