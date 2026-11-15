package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{12, 15, 11, 18, 21, 19, 14, 13}
	fmt.Println("temperatures (°C):", temps)
	fmt.Printf("averages over 3:  %.1f\n", MovingAverage(temps, 3))

	n, k := 100_000, 1440 // readings, one per minute; a 24-hour window
	big := make([]float64, n)
	for i := range big {
		big[i] = 10 + 8*math.Sin(2*math.Pi*float64(i)/1440)
	}
	Adds = 0
	MovingAverage(big, k)
	fmt.Printf("\n%d readings, window %d:\n  sliding:     %d additions\n  recomputing: %d additions\n", n, k, Adds, (n-k+1)*k)
}
