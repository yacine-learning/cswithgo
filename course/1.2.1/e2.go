package main

// Vec is a growable array. data is the backing array (its length is the capacity);
// only data[:n] is in use.
type Vec struct {
	data   []int
	n      int
	Copies int // elements copied when the array grew
}

func (v *Vec) Push(x int) {
	if v.n == len(v.data) {
		// hole 1: the new capacity is double the old one (1 when the old one is 0)
		newCap := 0
		_ = newCap // delete this line once you use newCap

		// hole 2: allocate an array of newCap elements, copy the n elements into it
		// (add 1 to v.Copies for each), then make it v.data

	}
	v.data[v.n] = x
	v.n++
}

func (v *Vec) Get(i int) int { return v.data[i] }
func (v *Vec) Len() int      { return v.n }
func (v *Vec) Cap() int      { return len(v.data) }
