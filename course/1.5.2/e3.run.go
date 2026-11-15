package main

import "fmt"

func main() {
	a := []Entry{{100, "A: Jupiter rises"}, {105, "A: meteor, north"}, {112, "A: Io transit begins"}, {130, "A: clouds"}}
	b := []Entry{{101, "B: Jupiter rises"}, {105, "B: meteor, north"}, {140, "B: Moon sets"}}
	for _, e := range Merge(a, b) {
		fmt.Println(e.T, e.Msg)
	}
}
