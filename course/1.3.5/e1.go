package main

// Du reads lines "PATH SIZE" (one per file; PATH has no spaces, SIZE is an integer)
// and returns one line "SIZE\tPATH" per directory, like `du -d maxDepth`:
//
//   - a directory's SIZE is the total of all files below it, at any depth;
//   - a top-level directory (the first part of a path) has depth 0, its subdirectories 1, …;
//     only directories with depth <= maxDepth are listed;
//   - order: post-order, like du: inside a directory, its subdirectories first, sorted by
//     name, then the directory itself; top-level directories sorted by name;
//   - a file at the top level (a PATH without "/") is in no directory and is not listed;
//   - a line that cannot be read gives the error "line N: cannot read \"...\"" (N from 1).
func Du(lines []string, maxDepth int) ([]string, error) {
	return nil, nil
}
