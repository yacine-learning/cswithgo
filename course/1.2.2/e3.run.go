package main

import "fmt"

func main() {
	darwin := `There is grandeur in this view of life, with its several powers, having been
originally breathed into a few forms or into one; and that, whilst this planet has gone
cycling on according to the fixed law of gravity, from so simple a beginning endless forms
most beautiful and most wonderful have been, and are being, evolved.`
	for i, c := range WordCounts(darwin) {
		if i == 8 {
			break
		}
		fmt.Printf("%-10s %d\n", c.Word, c.N)
	}
}
