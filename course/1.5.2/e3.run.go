package main

import "fmt"

func main() {
	podA := []Entry{{100, "a: start"}, {105, "a: request 1"}, {112, "a: request 2"}, {130, "a: shutdown"}}
	podB := []Entry{{101, "b: start"}, {105, "b: request 9"}, {140, "b: health check"}}
	for _, e := range Merge(podA, podB) {
		fmt.Println(e.T, e.Msg)
	}
}
