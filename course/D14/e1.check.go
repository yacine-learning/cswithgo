package main

import "fmt"

func main() {
	b := NewBloom(10000, 7)
	for i := 0; i < 1000; i++ {
		b.Add(fmt.Sprintf("word-%d", i))
	}
	try("never a false negative", func() string {
		for i := 0; i < 1000; i++ {
			if w := fmt.Sprintf("word-%d", i); !b.MayContain(w) {
				return fmt.Sprintf("%q was added, but MayContain(%q) returned false", w, w)
			}
		}
		return ""
	})
	try("an empty filter contains nothing", func() string {
		if NewBloom(100, 3).MayContain("go") {
			return `nothing was added, but MayContain("go") returned true`
		}
		return ""
	})
	try("few false positives", func() string {
		fp := 0
		for i := 0; i < 20000; i++ {
			if b.MayContain(fmt.Sprintf("other-%d", i)) {
				fp++
			}
		}
		// with 1 000 words, 10 000 bits and k = 7, theory says about 0.8 %
		if rate := 100 * float64(fp) / 20000; rate > 3 {
			return fmt.Sprintf("1 000 words, 10 000 bits, k = 7: %.1f %% of 20 000 words never added are reported present; expected about 1 %%", rate)
		}
		return ""
	})
	try("Add sets at most k bits", func() string {
		c := NewBloom(1000, 4)
		c.Add("gopher")
		set := 0
		for _, x := range c.bits {
			if x {
				set++
			}
		}
		if set == 0 || set > 4 {
			return fmt.Sprintf("after one Add with k = 4, %d bits are set", set)
		}
		return ""
	})
}

// try runs one check: an empty message is a pass, a panic is a failure.
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
