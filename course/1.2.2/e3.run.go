package main

import "fmt"

func main() {
	lovelace := `The Analytical Engine has no pretensions whatever to originate anything.
It can do whatever we know how to order it to perform. It can follow analysis;
but it has no power of anticipating any analytical relations or truths.`
	for i, c := range WordCounts(lovelace) {
		if i == 8 {
			break
		}
		fmt.Printf("%-10s %d\n", c.Word, c.N)
	}
}
