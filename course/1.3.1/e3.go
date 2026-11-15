package main

// Flatten turns a nested document into flat paths joined by dots:
// {"orbit": {"period": 365, "tilt": 23}} -> {"orbit.period": "365", "orbit.tilt": "23"}.
// Lists use the index of each element as its key: {"moons": ["Io", "Europa"]} -> {"moons.0": "Io", "moons.1": "Europa"}.
// Leaf values are written with fmt.Sprint.
func Flatten(doc map[string]any) map[string]string {
	out := map[string]string{}
	return out
}
