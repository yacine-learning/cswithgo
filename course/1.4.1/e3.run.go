package main

import "fmt"

func main() {
	pond := Graph{
		"heron":     {"frog", "fish"},
		"frog":      {"dragonfly", "snail"},
		"fish":      {"dragonfly", "algae"},
		"dragonfly": {"mosquito"},
		"mosquito":  {"algae"},
		"snail":     {"algae", "pondweed"},
		"algae":     nil,
		"pondweed":  nil,
	}
	fmt.Println("producers:      ", Producers(pond))
	fmt.Println("apex predators: ", ApexPredators(pond))
}
