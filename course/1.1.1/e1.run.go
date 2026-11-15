package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("n        comparisons   ×        time (best of 3)   ×")
	var prevSteps int
	var prevTime time.Duration
	for n := 2000; n <= 16000; n *= 2 {
		xs := make([]int, n)
		for i := range xs {
			xs[i] = i // all different: the worst case, every pair is compared
		}
		best := time.Duration(1 << 62)
		for try := 0; try < 3; try++ {
			Steps = 0
			start := time.Now()
			HasDuplicateSlow(xs)
			if d := time.Since(start); d < best {
				best = d
			}
		}
		rs, rt := "", ""
		if prevSteps > 0 {
			rs = fmt.Sprintf("×%.2f", float64(Steps)/float64(prevSteps))
			rt = fmt.Sprintf("×%.1f", float64(best)/float64(prevTime))
		}
		fmt.Printf("%-8d %-13d %-8s %-18s %s\n", n, Steps, rs, best.Round(time.Millisecond), rt)
		prevSteps, prevTime = Steps, best
	}
	fmt.Println("\nn doubles on each line. What do the comparisons do? And the time?")
}
