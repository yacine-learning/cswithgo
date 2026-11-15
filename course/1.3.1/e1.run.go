package main

import "fmt"

func main() {
	app := &Dir{Name: "app", Files: []File{{"README.md", 2}}, Subdirs: []*Dir{
		{Name: "src", Files: []File{{"main.go", 12}, {"util.go", 5}}, Subdirs: []*Dir{
			{Name: "api", Files: []File{{"api.go", 20}}},
		}},
		{Name: "config", Files: []File{{"app.yml", 1}}},
	}}
	fmt.Println("\ntotal:", DirSize(app, 0))
}
