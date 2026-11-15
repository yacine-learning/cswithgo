package main

import "fmt"

func main() {
	progs := []Process{{"edit", 5}, {"listf", 1}, {"mad", 4}, {"fap", 2}}
	fmt.Println("quantum 2:", RoundRobin(progs, 2))
	fmt.Println("quantum 10:", RoundRobin(progs, 10))
	fmt.Println("quantum 1:", RoundRobin(progs, 1))
}
