package main

// Sizes reads lines "PATH SIZE" (one per file; PATH has no spaces, SIZE is an integer)
// and returns one line "SIZE\tPATH" per folder:
//
//   - a folder's SIZE is the total of all files below it, at any depth;
//   - a top-level folder (the first part of a path) has depth 0, its subfolders 1, …;
//     only folders with depth <= maxDepth are listed;
//   - order: post-order: inside a folder, its subfolders first, sorted by
//     name, then the folder itself; top-level folders sorted by name;
//   - a file at the top level (a PATH without "/") is in no folder and is not listed;
//   - a line that cannot be read gives the error "line N: cannot read \"...\"" (N from 1).
func Sizes(lines []string, maxDepth int) ([]string, error) {
	return nil, nil
}
