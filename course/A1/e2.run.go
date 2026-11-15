package main

import "fmt"

func main() {
	fns := []struct {
		name string
		f    func(int) int
	}{{"A", stepsA}, {"B", stepsB}, {"C", stepsC}, {"D", stepsD}}
	for _, fn := range fns {
		fmt.Printf("%s  your answer: %-9s", fn.name, Answers[fn.name])
		prev := 0
		for n := 1000; n <= 8000; n *= 2 {
			s := fn.f(n)
			if prev > 0 {
				fmt.Printf("  n=%d: %d steps (×%.2f)", n, s, float64(s)/float64(prev))
			} else {
				fmt.Printf("  n=%d: %d steps", n, s)
			}
			prev = s
		}
		fmt.Println()
	}
}
