package main

// Job needs Time units of work.
type Job struct {
	Name string
	Time int
}

// RoundRobin runs the jobs in turn, each for at most quantum units, and
// returns their names in the order they finish.
func RoundRobin(jobs []Job, quantum int) []string {
	q := append([]Job(nil), jobs...) // a copy: do not change the caller's slice
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
