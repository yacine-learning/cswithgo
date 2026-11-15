package main

// Bloom is a set that can answer "definitely not" or "probably yes", using one bit per slot.
type Bloom struct {
	bits []bool
	k    int // how many slots each value sets
}

// NewBloom makes a filter of m bits where each value sets k of them.
func NewBloom(m, k int) *Bloom { return &Bloom{bits: make([]bool, m), k: k} }

// slots returns the k positions of s in the filter. One 64-bit hash is cut in two halves,
// h1 and h2, and the i-th position is h1 + i*h2 (Kirsch and Mitzenmacher, 2006): as good
// as k separate hash functions, for the price of one.
func (b *Bloom) slots(s string) []int {
	v := fnv1a(s)
	h1, h2 := v&0xffffffff, v>>32|1
	out := make([]int, b.k)
	for i := range out {
		out[i] = int((h1 + uint64(i)*h2) % uint64(len(b.bits)))
	}
	return out
}

// fnv1a is the 64-bit FNV-1a hash (Fowler, Noll and Vo): for each byte, xor it in, then
// multiply by a large prime. The same as Go's hash/fnv New64a.
func fnv1a(s string) uint64 {
	h := uint64(14695981039346656037)
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}

// Add puts s in the filter.
func (b *Bloom) Add(s string) {
	for _, p := range b.slots(s) {
		b.bits[p] = true
	}
}

// MayContain answers false only when s was certainly never added.
func (b *Bloom) MayContain(s string) bool {
	for _, p := range b.slots(s) {
		if !b.bits[p] {
			return false
		}
	}
	return true
}
