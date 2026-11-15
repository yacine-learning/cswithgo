package main

import "fmt"

func main() {
	try("counts and order", func() string {
		got := StatusCounts([]string{
			"GET /a 200 1ms", "GET /b 500 2ms", "GET /c 200 1ms",
			"POST /d 401 1ms", "GET /e 500 9ms", "GET /f 200 1ms",
		})
		want := []Count{{"200", 3}, {"500", 2}, {"401", 1}}
		if fmt.Sprint(got) != fmt.Sprint(want) {
			return fmt.Sprintf("got %v, expected %v", got, want)
		}
		return ""
	})
	try("ties: smaller code first", func() string {
		got := StatusCounts([]string{"GET /a 503 1ms", "GET /b 404 1ms", "GET /c 200 1ms"})
		want := []Count{{"200", 1}, {"404", 1}, {"503", 1}}
		if fmt.Sprint(got) != fmt.Sprint(want) {
			return fmt.Sprintf("got %v, expected %v", got, want)
		}
		return ""
	})
	try("no lines", func() string {
		if got := StatusCounts(nil); len(got) != 0 {
			return fmt.Sprintf("got %v for no lines, expected nothing", got)
		}
		return ""
	})
	try("the same result every run (no map order leaks out)", func() string {
		lines := []string{}
		for i := 0; i < 50; i++ {
			lines = append(lines, fmt.Sprintf("GET /x %d 1ms", 200+i%7))
		}
		first := fmt.Sprint(StatusCounts(lines))
		for i := 0; i < 20; i++ {
			if s := fmt.Sprint(StatusCounts(lines)); s != first {
				return "two calls on the same lines gave different orders:\n  " + first + "\n  " + s
			}
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
