package main

// Plan groups packages into build rounds.
//
// deps[p] lists the packages that p imports: each of them must be built in an earlier
// round than p. A package that appears only as an import is a package too.
//
//   - round 0 holds every package that imports nothing; round k holds every package whose
//     imports are all built in rounds before k (each package as early as possible);
//   - the packages of each round are sorted by name;
//   - when some packages can never be built (they are on a cycle, or import one),
//     Plan returns the error "cannot order: a, b, c": all those packages, sorted.
func Plan(deps map[string][]string) ([][]string, error) {
	return nil, nil
}
