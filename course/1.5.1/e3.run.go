package main

import "fmt"

func main() {
	calls := 0
	firstBroken := 6123
	isBad := func(b int) bool {
		calls++
		bad := b >= firstBroken
		fmt.Printf("  test build %5d: %v\n", b, map[bool]string{true: "bad", false: "good"}[bad])
		return bad
	}
	got := FirstBad(10000, isBad)
	fmt.Printf("first bad build: %d, found with %d test runs out of 10 000 builds\n", got, calls)
}
