package main

import (
	"fmt"
	"strings"
)

func main() {
	traces := map[string]string{
		"a good trace": `10 start main
12 start parse
15 start lex
41 end lex
44 end parse
45 start eval
60 end eval
61 end main`,
		"ends in the wrong order": `1 start parse
2 start lex
3 end parse`,
		"left open": `1 start parse
2 start lex
3 end lex`,
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
