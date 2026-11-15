package main

import "fmt"

func main() {
	log := []string{
		"GET /api/users 200 12ms",
		"GET /api/users 200 9ms",
		"POST /api/login 401 3ms",
		"GET /healthz 200 1ms",
		"GET /api/orders 500 230ms",
		"POST /api/login 200 15ms",
		"GET /api/orders 500 210ms",
	}
	for _, c := range StatusCounts(log) {
		fmt.Printf("%s  %d\n", c.Code, c.N)
	}
}
