package main

import "fmt"

func main() {
	pods := [][]Entry{
		{{100, "pod-0 GET /cart"}, {103, "pod-0 200"}, {110, "pod-0 GET /pay"}},
		{{101, "pod-1 GET /cart"}, {103, "pod-1 500"}},
		{{99, "pod-2 started"}, {120, "pod-2 GET /health"}},
	}
	for _, e := range MergeAll(pods) {
		fmt.Println(e.T, e.Msg)
	}
}
