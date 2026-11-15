package main

// WithFlag returns the arguments of base followed by flag.
// base must not change, and each result must be independent of the others.
func WithFlag(base []string, flag string) []string {
	return append(base, flag)
}
