package main

import "fmt"

func main() {
	// each thread's trace buffer, sorted by time (microseconds)
	threads := [][]Entry{
		{{100, "thread 0: main starts"}, {103, "thread 0: spawns 1 and 2"}, {110, "thread 0: waits"}},
		{{101, "thread 1: idle"}, {103, "thread 1: reads input"}, {118, "thread 1: done"}},
		{{99, "thread 2: idle"}, {104, "thread 2: sorts half"}, {120, "thread 2: done"}},
	}
	for _, e := range MergeAll(threads) {
		fmt.Println(e.T, e.Msg)
	}
}
