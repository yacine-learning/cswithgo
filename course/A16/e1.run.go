package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(" n   solutions   every board with one queen per row (n^n)   time")
	for n := 1; n <= 12; n++ {
		t := time.Now()
		s := Queens(n)
		fmt.Printf("%2d   %9d   %42.0f   %s\n", n, s, math.Pow(float64(n), float64(n)), time.Since(t).Round(time.Microsecond))
	}
}
