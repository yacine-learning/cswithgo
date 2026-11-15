package main

// Graph maps each service to the services it depends on.
type Graph map[string][]string

// DeployFirst returns the services that depend on nothing, sorted.
func DeployFirst(g Graph) []string {
	return nil
}

// SafeToRemove returns the services that no other service depends on, sorted.
func SafeToRemove(g Graph) []string {
	return nil
}
