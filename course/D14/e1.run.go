package main

import "fmt"

func main() {
	// 1 000 words in a filter of 10 000 bits (1 250 bytes once packed 8 to a byte; here a
	// []bool keeps it simple), then 100 000 words that were never added: how often does it lie?
	n, m := 1000, 10000
	fmt.Printf("%d words, %d bits\n\n k   bits set   false positives\n", n, m)
	for k := 1; k <= 12; k++ {
		b := NewBloom(m, k)
		for i := 0; i < n; i++ {
			b.Add(fmt.Sprintf("word-%d", i))
		}
		set := 0
		for _, x := range b.bits {
			if x {
				set++
			}
		}
		fp := 0
		for i := 0; i < 100000; i++ {
			if b.MayContain(fmt.Sprintf("other-%d", i)) {
				fp++
			}
		}
		fmt.Printf("%2d   %5.1f %%     %6.2f %%\n", k, 100*float64(set)/float64(m), 100*float64(fp)/100000)
	}
}
