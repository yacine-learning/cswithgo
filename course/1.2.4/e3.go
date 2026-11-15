package main

// Job needs Time units of work.
type Job struct {
	Name string
	Time int
}

// RoundRobin runs the jobs in turn, each for at most quantum units, and
// returns their names in the order they finish.
func RoundRobin(jobs []Job, quantum int) []string {
	return nil
}
