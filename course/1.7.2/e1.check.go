package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	try("three threads, a tie at 103", func() string {
		threads := [][]Entry{
			{{100, "a1"}, {103, "a2"}, {110, "a3"}},
			{{101, "b1"}, {103, "b2"}},
			{{99, "c1"}, {120, "c2"}},
		}
		want := "[{99 c1} {100 a1} {101 b1} {103 a2} {103 b2} {110 a3} {120 c2}]"
		if got := fmt.Sprint(MergeAll(threads)); got != want {
			return "got " + got + "\n      expected " + want
		}
		return ""
	})
	try("equal times inside one thread keep their order", func() string {
		threads := [][]Entry{{{5, "x1"}, {5, "x2"}}, {{5, "y1"}}}
		if got := fmt.Sprint(MergeAll(threads)); got != "[{5 x1} {5 x2} {5 y1}]" {
			return "got " + got + ", expected [{5 x1} {5 x2} {5 y1}]"
		}
		return ""
	})
	try("empty threads and no threads", func() string {
		if got := MergeAll([][]Entry{nil, {{1, "a"}}, {}}); fmt.Sprint(got) != "[{1 a}]" {
			return fmt.Sprintf("MergeAll(nil, [{1 a}], []) = %v", got)
		}
		if got := MergeAll(nil); len(got) != 0 {
			return fmt.Sprintf("MergeAll(nil) = %v", got)
		}
		return ""
	})
	try("300 threads × 1 000 events", func() string {
		rng := rand.New(rand.NewSource(11))
		threads := make([][]Entry, 300)
		type tagged struct {
			e         Entry
			thread, i int
		}
		var all []tagged
		for p := range threads {
			t := 0
			for i := 0; i < 1000; i++ {
				t += rng.Intn(5)
				e := Entry{t, fmt.Sprintf("%d/%d", p, i)}
				threads[p] = append(threads[p], e)
				all = append(all, tagged{e, p, i})
			}
		}
		sort.Slice(all, func(a, b int) bool {
			if all[a].e.T != all[b].e.T {
				return all[a].e.T < all[b].e.T
			}
			if all[a].thread != all[b].thread {
				return all[a].thread < all[b].thread
			}
			return all[a].i < all[b].i
		})
		got := MergeAll(threads)
		if len(got) != len(all) {
			return fmt.Sprintf("%d entries out, %d in", len(got), len(all))
		}
		for i := range all {
			if got[i] != all[i].e {
				return fmt.Sprintf("entry %d is %v, expected %v", i, got[i], all[i].e)
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
