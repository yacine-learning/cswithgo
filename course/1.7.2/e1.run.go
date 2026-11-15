package main

import "fmt"

func main() {
	stations := [][]Entry{
		{{100, "Basel 4.1°C"}, {103, "Basel 3.9°C"}, {110, "Basel 3.2°C, rain"}},
		{{101, "Geneva 6.0°C"}, {103, "Geneva 5.4°C, wind 40 km/h"}},
		{{99, "Säntis -8.5°C"}, {120, "Säntis -9.1°C, snow"}},
	}
	for _, e := range MergeAll(stations) {
		fmt.Println(e.T, e.Msg)
	}
}
