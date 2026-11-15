package main

import (
	"fmt"
	"strings"
)

func main() {
	traces := map[string]string{
		"a good trace": `10 start pipeline
12 start build
15 start test
41 end test
44 end build
45 start deploy
60 end deploy
61 end pipeline`,
		"ends in the wrong order": `1 start build
2 start test
3 end build`,
		"left open": `1 start build
2 start test
3 end test`,
	}
	for _, name := range []string{"a good trace", "ends in the wrong order", "left open"} {
		r, err := Analyze(strings.Split(traces[name], "\n"))
		fmt.Println(name + ":")
		if err != nil {
			fmt.Println("  error:", err)
			continue
		}
		fmt.Println("  durations:", r.Durations)
		fmt.Println("  max depth:", r.MaxDepth)
		fmt.Println("  slowest:  ", r.Slowest)
	}
}
