package main

import "fmt"

func main() {
	counts := map[string]int{
		"connection refused": 412, "timeout": 977, "401 unauthorized": 58,
		"disk full": 3, "502 bad gateway": 412, "DNS lookup failed": 120,
	}
	fmt.Println("top 3:", TopK(counts, 3))
}
