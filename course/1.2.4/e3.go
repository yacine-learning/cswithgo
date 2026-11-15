package main

// Process needs Time units of processor time.
type Process struct {
	Name string
	Time int
}

// RoundRobin runs the processes in turn, each for at most quantum units, and
// returns their names in the order they finish.
func RoundRobin(procs []Process, quantum int) []string {
	return nil
}
