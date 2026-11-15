package main

import "fmt"

func expect(jobs []Job, q int, want string) func() string {
	return func() string {
		got := fmt.Sprint(RoundRobin(jobs, q))
		if got != want {
			return fmt.Sprintf("RoundRobin(%v, %d) = %s, expected %s", jobs, q, got, want)
		}
		return ""
	}
}

func main() {
	jobs := []Job{{"build", 5}, {"lint", 1}, {"test", 4}, {"docs", 2}}
	try("quantum 2", expect(jobs, 2, "[lint docs test build]"))
	try("a quantum bigger than every job: plain first come, first served", expect(jobs, 10, "[build lint test docs]"))
	try("quantum 1", expect(jobs, 1, "[lint docs test build]"))
	try("a job that needs exactly one quantum finishes in its turn", expect([]Job{{"a", 3}, {"b", 3}}, 3, "[a b]"))
	try("no jobs", expect(nil, 2, "[]"))
	try("the caller's slice is not changed", func() string {
		js := []Job{{"a", 2}, {"b", 1}}
		RoundRobin(js, 1)
		if js[0].Time != 2 || js[1].Time != 1 {
			return fmt.Sprintf("after the call, the jobs passed in became %v", js)
		}
		return ""
	})
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
