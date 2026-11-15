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
		return same(Flatten(map[string]any{"name": "Earth", "moons": 1}), map[string]string{"name": "Earth", "moons": "1"})
	})
	try("one level of nesting", func() string {
		return same(Flatten(map[string]any{"orbit": map[string]any{"period": 365, "tilt": 23}}),
			map[string]string{"orbit.period": "365", "orbit.tilt": "23"})
	})
	try("deep nesting", func() string {
		return same(Flatten(map[string]any{"a": map[string]any{"b": map[string]any{"c": map[string]any{"d": true}}}}),
			map[string]string{"a.b.c.d": "true"})
	})
	try("lists use indexes", func() string {
		return same(Flatten(map[string]any{"moons": []any{"Io", "Europa"}}), map[string]string{"moons.0": "Io", "moons.1": "Europa"})
	})
	try("maps inside lists", func() string {
		return same(Flatten(map[string]any{"planets": []any{map[string]any{"name": "Mercury"}, map[string]any{"name": "Venus"}}}),
			map[string]string{"planets.0.name": "Mercury", "planets.1.name": "Venus"})
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
