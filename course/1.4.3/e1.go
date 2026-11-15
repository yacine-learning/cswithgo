package main

// Plan groups services into deployment waves.
//
// deps[s] lists the services s depends on: each of them must be in an earlier wave than s.
// A service that appears only as a dependency is a service too.
//
//   - wave 0 holds every service with no dependency; wave k holds every service whose
//     dependencies are all in waves before k (each service as early as possible);
//   - the services of each wave are sorted by name;
//   - when some services can never be deployed (they are on a cycle, or depend on one),
//     Plan returns the error "cannot order: a, b, c": all those services, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	return nil, nil
}
