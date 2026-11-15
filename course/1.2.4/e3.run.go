package main

import "fmt"

func main() {
	jobs := []Job{{"build", 5}, {"lint", 1}, {"test", 4}, {"docs", 2}}
	fmt.Println("quantum 2:", RoundRobin(jobs, 2))
	fmt.Println("quantum 10:", RoundRobin(jobs, 10))
	fmt.Println("quantum 1:", RoundRobin(jobs, 1))
}
