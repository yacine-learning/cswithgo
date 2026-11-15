package main

// Graph maps each package to the packages it imports.
type Graph map[string][]string

// Leaves returns the packages that import nothing, sorted.
func Leaves(g Graph) []string {
	return nil
}

// Roots returns the packages that no other package imports, sorted.
func Roots(g Graph) []string {
	return nil
}
