package main

import "fmt"

func main() {
	ts := []int{0, 5, 30, 31, 31, 32, 33, 60, 61, 300, 301, 302, 303, 304, 305, 306, 900}
	try("Count in a range", func() string {
		for _, c := range []struct{ from, to, want int }{
			{30, 60, 5}, {0, 1, 1}, {31, 32, 2}, {61, 300, 1}, {-10, 1000, 17}, {901, 2000, 0}, {50, 50, 0}, {60, 30, 0},
		} {
			if got := Count(ts, c.from, c.to); got != c.want {
				return fmt.Sprintf("Count(ts, %d, %d) = %d, expected %d", c.from, c.to, got, c.want)
			}
		}
		return ""
	})
	try("Busiest window", func() string {
		if s, n := Busiest(ts, 10); s != 300 || n != 7 {
			return fmt.Sprintf("Busiest(ts, 10) = %d, %d; expected 300, 7", s, n)
		}
		if s, n := Busiest(ts, 2); s != 30 || n != 3 {
			return fmt.Sprintf("Busiest(ts, 2) = %d, %d; expected 30, 3 (30 31 31; the earliest of the ties)", s, n)
		}
		if s, n := Busiest(ts, 10000); s != 0 || n != 17 {
			return fmt.Sprintf("Busiest(ts, 10000) = %d, %d; expected 0, 17", s, n)
		}
		if s, n := Busiest(nil, 10); s != 0 || n != 0 {
			return fmt.Sprintf("Busiest(nil, 10) = %d, %d; expected 0, 0", s, n)
		}
		return ""
	})
	try("Silences", func() string {
		if got := fmt.Sprint(Silences(ts, 120)); got != "[[61 300] [306 900]]" {
			return "Silences(ts, 120) = " + got + ", expected [[61 300] [306 900]]"
		}
		if got := Silences(ts, 1000); len(got) != 0 {
			return fmt.Sprintf("Silences(ts, 1000) = %v, expected none", got)
		}
		if got := fmt.Sprint(Silences([]int{1, 5}, 4)); got != "[[1 5]]" {
			return "Silences([1 5], 4) = " + got + ", expected [[1 5]] (a gap of exactly min counts)"
		}
		return ""
	})
	try("Busiest on 300 000 errors (O(n))", func() string {
		big := make([]int, 300000)
		for i := range big {
			big[i] = i / 3
		}
		if s, n := Busiest(big, 1000); n != 3000 || s != 0 {
			return fmt.Sprintf("Busiest = %d, %d; expected 0, 3000", s, n)
		}
		return ""
	})
	try("50 000 Count queries on 500 000 errors", func() string {
		big := make([]int, 500000)
		for i := range big {
			big[i] = i * 2
		}
		for q := 0; q < 50000; q++ {
			from := (q * 37) % 900000
			if got := Count(big, from, from+1000); got != 500 {
				return fmt.Sprintf("Count(big, %d, %d) = %d, expected 500", from, from+1000, got)
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
