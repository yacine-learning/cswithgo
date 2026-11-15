package main

// Dir is a directory (or a file, when it has no Children).
type Dir struct {
	Name     string
	Children []*Dir
}

// CreateOrder returns every path, each directory before what it contains.
// The root's path is its name; a child's path is its parent's path + "/" + its name.
func CreateOrder(root *Dir) []string {
	return nil
}

// DeleteOrder returns every path, everything a directory contains before the directory.
func DeleteOrder(root *Dir) []string {
	return nil
}
