package main

// WithFlag returns the arguments of base followed by flag.
// base must not change, and each result must be independent of the others.
func WithFlag(base []string, flag string) []string {
	out := make([]string, len(base), len(base)+1)
	copy(out, base)
	return append(out, flag)
}
