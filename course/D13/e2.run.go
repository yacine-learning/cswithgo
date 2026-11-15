package main

import "fmt"

func main() {
	// six machines in a lab, and the cables we could lay (cost in metres)
	links := []Link{
		{0, 1, 7}, {0, 3, 5}, {1, 2, 8}, {1, 3, 9}, {1, 4, 7}, {2, 4, 5},
		{3, 4, 15}, {3, 5, 6}, {4, 5, 8}, {4, 0, 12},
	}
	cost, chosen, ok := Cheapest(6, links)
	fmt.Println("all connected:", ok)
	for _, l := range chosen {
		fmt.Printf("  lay %d-%d (%d m)\n", l.A, l.B, l.Cost)
	}
	all := 0
	for _, l := range links {
		all += l.Cost
	}
	fmt.Printf("total: %d m, instead of %d m for every cable\n", cost, all)

	_, _, ok = Cheapest(4, []Link{{0, 1, 1}, {2, 3, 1}})
	fmt.Println("\ntwo islands, 0-1 and 2-3, connected:", ok)
}
