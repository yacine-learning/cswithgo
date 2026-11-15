package main

// Count is how many log lines had one status code.
type Count struct {
	Code string
	N    int
}

// StatusCounts counts the status codes of access-log lines such as
// "GET /api/users 200 12ms", most frequent first, then by code.
func StatusCounts(lines []string) []Count {
	return nil
}
