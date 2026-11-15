package main

import "fmt"

func show(name string, deps map[string][]string) {
	fmt.Println(name + ":")
	waves, err := Plan(deps)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	for i, w := range waves {
		fmt.Printf("  wave %d: %v\n", i, w)
	}
}

func main() {
	show("a release", map[string][]string{
		"web":    {"api"},
		"api":    {"db", "cache", "queue"},
		"worker": {"db", "queue"},
	})
	show("with a cycle", map[string][]string{
		"web":   {"api"},
		"api":   {"auth"},
		"auth":  {"users"},
		"users": {"api"},
		"cron":  {"db"},
	})
}
