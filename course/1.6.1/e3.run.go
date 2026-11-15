package main

import "fmt"

func main() {
	running := []Interval{{120, 180}, {0, 30}, {150, 200}, {30, 45}, {300, 360}, {170, 190}}
	fmt.Println("threads ran:", running)
	fmt.Println("the processor was busy:", MergeIntervals(running))
}
