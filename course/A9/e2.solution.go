package main

// Dir is a directory (or a file, when it has no Children).
type Dir struct {
	Name     string
	Children []*Dir
}

// CreateOrder returns every path, each directory before what it contains.
// The root's path is its name; a child's path is its parent's path + "/" + its name.
func CreateOrder(root *Dir) []string {
	var out []string
	var pre func(d *Dir, path string)
	pre = func(d *Dir, path string) {
		out = append(out, path)
		for _, c := range d.Children {
			pre(c, path+"/"+c.Name)
		}
	}
	pre(root, root.Name)
	return out
}

// DeleteOrder returns every path, everything a directory contains before the directory.
func DeleteOrder(root *Dir) []string {
	var out []string
	var post func(d *Dir, path string)
	post = func(d *Dir, path string) {
		for _, c := range d.Children {
			post(c, path+"/"+c.Name)
		}
		out = append(out, path)
	}
	post(root, root.Name)
	return out
}
