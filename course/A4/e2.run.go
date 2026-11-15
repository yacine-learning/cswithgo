package main

import "fmt"

func main() {
	cpu0 := []Entry{{100, "cpu0: main starts"}, {105, "cpu0: go worker()"}, {112, "cpu0: waits on channel"}, {130, "cpu0: receives result"}}
	cpu1 := []Entry{{101, "cpu1: idle"}, {105, "cpu1: worker starts"}, {128, "cpu1: sends result"}, {140, "cpu1: idle"}}
	for _, e := range Merge(cpu0, cpu1) {
		fmt.Println(e.T, e.Msg)
	}
}
