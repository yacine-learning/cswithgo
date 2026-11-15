package main

import "fmt"

func main() {
	observed := []Interval{{120, 180}, {0, 30}, {150, 200}, {30, 45}, {300, 360}, {170, 190}}
	fmt.Println("observed:", observed)
	fmt.Println("merged:", MergeIntervals(observed))
}
