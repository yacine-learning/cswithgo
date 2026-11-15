package main

import "fmt"

func main() {
	g := Graph{
		"web":    {"api"},
		"api":    {"db", "cache", "queue"},
		"worker": {"db", "queue"},
		"admin":  {"db"},
		"db":     nil,
		"cache":  nil,
		"queue":  nil,
	}
	fmt.Println("deploy first:  ", DeployFirst(g))
	fmt.Println("safe to remove:", SafeToRemove(g))
}
