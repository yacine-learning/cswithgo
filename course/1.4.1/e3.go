package main

// Graph maps each species to the species it eats.
type Graph map[string][]string

// Producers returns the species that eat nothing, sorted.
func Producers(g Graph) []string {
	return nil
}

// ApexPredators returns the species that no other species eats, sorted.
func ApexPredators(g Graph) []string {
	return nil
}
