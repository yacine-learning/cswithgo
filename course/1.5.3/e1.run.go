package main

import "fmt"

func main() {
	lat := []float64{12, 15, 11, 40, 38, 13, 12, 14}
	fmt.Println("latencies (ms):", lat)
	fmt.Printf("averages over 3: %.1f\n", MovingAverage(lat, 3))

	n, k := 100_000, 1000
	big := make([]float64, n)
	for i := range big {
		big[i] = float64(i % 97)
	}
	Adds = 0
	MovingAverage(big, k)
	fmt.Printf("\n%d values, window %d:\n  sliding:     %d additions\n  recomputing: %d additions\n", n, k, Adds, (n-k+1)*k)
}
