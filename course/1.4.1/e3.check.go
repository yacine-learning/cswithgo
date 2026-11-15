package main

import "fmt"

func main() {
	g := Graph{
		"heron": {"frog", "fish"}, "frog": {"dragonfly", "snail"}, "fish": {"dragonfly", "algae"},
		"otter": {"fish"}, "dragonfly": {"mosquito"}, "mosquito": {"algae"}, "snail": {"algae", "pondweed"},
		"algae": nil, "pondweed": nil,
	}
	try("producers (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(Producers(g)); got != "[algae pondweed]" {
				return "Producers = " + got + ", expected [algae pondweed]"
			}
		}
		return ""
	})
	try("apex predators (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(ApexPredators(g)); got != "[heron otter]" {
				return "ApexPredators = " + got + ", expected [heron otter]"
			}
		}
		return ""
	})
	try("a lone species is both", func() string {
		g := Graph{"lichen": nil}
		if fmt.Sprint(Producers(g)) != "[lichen]" || fmt.Sprint(ApexPredators(g)) != "[lichen]" {
			return fmt.Sprintf("Producers = %v, ApexPredators = %v, expected [lichen] for both", Producers(g), ApexPredators(g))
		}
		return ""
	})
	try("a chain hawk -> snake -> mouse -> grass", func() string {
		g := Graph{"hawk": {"snake"}, "snake": {"mouse"}, "mouse": {"grass"}, "grass": nil}
		if fmt.Sprint(Producers(g)) != "[grass]" || fmt.Sprint(ApexPredators(g)) != "[hawk]" {
			return fmt.Sprintf("Producers = %v, ApexPredators = %v; expected [grass] and [hawk]", Producers(g), ApexPredators(g))
		}
		return ""
	})
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
