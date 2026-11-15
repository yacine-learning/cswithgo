package main

import "fmt"

func main() {
	try("a small network", func() string {
		links := []Link{{0, 1, 4}, {1, 2, 1}, {0, 2, 3}, {2, 3, 2}, {1, 3, 5}}
		cost, chosen, ok := Cheapest(4, links)
		if !ok {
			return "4 machines that can all be connected: ok = false"
		}
		if cost != 6 {
			return fmt.Sprintf("links 0-1:4 1-2:1 0-2:3 2-3:2 1-3:5: cost %d, expected 6 (1-2, 2-3, 0-2)", cost)
		}
		if len(chosen) != 3 {
			return fmt.Sprintf("connecting 4 machines needs 3 links, %d were chosen", len(chosen))
		}
		return ""
	})
	try("a loop is never closed", func() string {
		links := []Link{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}, {2, 3, 10}}
		cost, chosen, ok := Cheapest(4, links)
		if !ok || cost != 12 || len(chosen) != 3 {
			return fmt.Sprintf("triangle 0-1-2 of cost 1 each, then 2-3 of cost 10: got cost %d, %d links, ok %v; expected 12, 3 links, true", cost, len(chosen), ok)
		}
		return ""
	})
	try("unreachable machine", func() string {
		_, _, ok := Cheapest(4, []Link{{0, 1, 1}, {1, 2, 1}})
		if ok {
			return "machine 3 has no link at all, but ok = true"
		}
		return ""
	})
	try("one machine needs nothing", func() string {
		cost, chosen, ok := Cheapest(1, nil)
		if !ok || cost != 0 || len(chosen) != 0 {
			return fmt.Sprintf("Cheapest(1, nil) = %d, %v, %v; expected 0, no links, true", cost, chosen, ok)
		}
		return ""
	})
	try("the input is not reordered", func() string {
		links := []Link{{0, 1, 9}, {1, 2, 1}, {0, 2, 5}}
		Cheapest(3, links)
		if links[0].Cost != 9 || links[1].Cost != 1 {
			return "Cheapest sorted the caller's slice; work on a copy"
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
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
