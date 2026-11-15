package main

import "fmt"

func main() {
	ids := []string{"u1", "u2", "u1", "u3", "u4", "u5", "u3", "u6"}
	s, n := LongestUnique(ids)
	fmt.Println("ids:", ids)
	fmt.Printf("longest run without a repeat: %v (start %d, length %d)\n", ids[s:s+n], s, n)
}
