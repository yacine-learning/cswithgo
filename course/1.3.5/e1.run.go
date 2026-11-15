package main

import "fmt"

func main() {
	files := []string{
		"app/README.md 2",
		"app/src/main.go 12",
		"app/src/api/api.go 20",
		"app/config/app.yml 1",
		"logs/2026-11-14.log 300",
		"logs/archive/2026-10.log.gz 900",
	}
	for _, d := range []int{0, 1, 2} {
		out, err := Du(files, d)
		fmt.Printf("$ du -d %d\n", d)
		if err != nil {
			fmt.Println("error:", err)
		}
		for _, l := range out {
			fmt.Println(l)
		}
		fmt.Println()
	}
}
