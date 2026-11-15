package main

// Process needs Time units of processor time.
type Process struct {
	Name string
	Time int
}

// RoundRobin runs the processes in turn, each for at most quantum units, and
// returns their names in the order they finish.
func RoundRobin(procs []Process, quantum int) []string {
	q := append([]Process(nil), procs...) // a copy: do not change the caller's slice
	var done []string
	for len(q) > 0 {
		j := q[0]
		q = q[1:]
		j.Time -= min(quantum, j.Time)
		if j.Time == 0 {
			done = append(done, j.Name)
		} else {
			q = append(q, j)
		}
	}
	return done
}
