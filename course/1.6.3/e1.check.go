package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	try("the example: 8 lines, memory 3", func() string {
		s, runs, rounds := ExternalSort([]string{"g", "c", "a", "h", "e", "b", "f", "d"}, 3)
		if fmt.Sprint(s) != "[a b c d e f g h]" || runs != 3 || rounds != 2 {
			return fmt.Sprintf("got %v, %d runs, %d rounds; expected [a b c d e f g h], 3 runs, 2 rounds", s, runs, rounds)
		}
		return ""
	})
	try("everything fits in memory: one run, no merge", func() string {
		s, runs, rounds := ExternalSort([]string{"b", "a", "c"}, 10)
		if fmt.Sprint(s) != "[a b c]" || runs != 1 || rounds != 0 {
			return fmt.Sprintf("got %v, %d runs, %d rounds; expected [a b c], 1 run, 0 rounds", s, runs, rounds)
		}
		return ""
	})
	try("memory 1: n runs", func() string {
		s, runs, rounds := ExternalSort([]string{"d", "c", "b", "a", "e"}, 1)
		if fmt.Sprint(s) != "[a b c d e]" || runs != 5 || rounds != 3 {
			return fmt.Sprintf("got %v, %d runs, %d rounds; expected [a b c d e], 5 runs, 3 rounds", s, runs, rounds)
		}
		return ""
	})
	try("no lines", func() string {
		s, runs, rounds := ExternalSort(nil, 4)
		if len(s) != 0 || runs != 0 || rounds != 0 {
			return fmt.Sprintf("got %v, %d, %d; expected nothing, 0, 0", s, runs, rounds)
		}
		return ""
	})
	try("the input is not changed", func() string {
		in := []string{"z", "y", "x", "w"}
		ExternalSort(in, 2)
		if fmt.Sprint(in) != "[z y x w]" {
			return fmt.Sprintf("the input became %v", in)
		}
		return ""
	})
	try("50 000 lines, memory 700", func() string {
		rng := rand.New(rand.NewSource(5))
		in := make([]string, 50000)
		for i := range in {
			in[i] = fmt.Sprintf("%08d", rng.Intn(100000000))
		}
		want := slices.Clone(in)
		slices.Sort(want)
		s, runs, rounds := ExternalSort(in, 700)
		if !slices.Equal(s, want) {
			return "the result is not the sorted input"
		}
		if runs != 72 || rounds != 7 {
			return fmt.Sprintf("%d runs, %d rounds; expected 72 runs (50 000 / 700, rounded up) and 7 rounds", runs, rounds)
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
