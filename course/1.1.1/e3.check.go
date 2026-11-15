package main

import "fmt"

func main() {
	want := map[string]Growth{"A": Logarithmic, "B": Linear, "C": Quadratic, "D": Constant}
	for _, k := range []string{"A", "B", "C", "D"} {
		got := Answers[k]
		switch {
		case got == want[k]:
			fmt.Printf("PASS %s is %s\n", k, got)
		case got == Unknown || got == "":
			fmt.Printf("FAIL %s: no answer yet\n", k)
		default:
			fmt.Printf("FAIL %s: you said %s; Run doubles n: compare the ratio with the table in the lesson\n", k, got)
		}
	}
}
