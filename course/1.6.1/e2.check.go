package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	try("merge two sorted halves", func() string {
		got := merge([]Rec{{1, "a"}, {4, "b"}, {9, "c"}}, []Rec{{2, "d"}, {3, "e"}, {10, "f"}})
		if fmt.Sprint(got) != "[{1 a} {2 d} {3 e} {4 b} {9 c} {10 f}]" {
			return fmt.Sprintf("merge = %v", got)
		}
		return ""
	})
	try("merge with an empty side", func() string {
		if got := merge(nil, []Rec{{1, "x"}}); fmt.Sprint(got) != "[{1 x}]" {
			return fmt.Sprintf("merge(nil, [{1 x}]) = %v", got)
		}
		if got := merge([]Rec{{1, "x"}}, nil); fmt.Sprint(got) != "[{1 x}]" {
			return fmt.Sprintf("merge([{1 x}], nil) = %v", got)
		}
		return ""
	})
	try("stable: equal keys keep their order", func() string {
		got := MergeSort([]Rec{{3, "api-1"}, {1, "web-1"}, {3, "api-2"}, {2, "db-1"}, {1, "web-2"}, {3, "api-3"}})
		want := "[{1 web-1} {1 web-2} {2 db-1} {3 api-1} {3 api-2} {3 api-3}]"
		if fmt.Sprint(got) != want {
			return fmt.Sprintf("got %v\n      expected %s", got, want)
		}
		return ""
	})
	try("10 000 random records", func() string {
		rng := rand.New(rand.NewSource(7))
		in := make([]Rec, 10000)
		for i := range in {
			in[i] = Rec{rng.Intn(100), fmt.Sprint(i)}
		}
		want := append([]Rec(nil), in...)
		sort.SliceStable(want, func(i, j int) bool { return want[i].Key < want[j].Key })
		got := MergeSort(in)
		if fmt.Sprint(got) != fmt.Sprint(want) {
			return "the result differs from a stable sort of the same records"
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
