package main

// Vec is a growable array, built the way Go's append works underneath.
// data is the backing array: its length is the capacity. Only data[:n] is in use.
type Vec struct {
	data   []int
	n      int
	Copies int // elements copied when the array grew
}

func (v *Vec) Push(x int) {
	if v.n == len(v.data) { // full: allocate a bigger array, copy everything over
		newCap := 2 * len(v.data)
		if newCap == 0 {
			newCap = 1
		}
		bigger := make([]int, newCap)
		for i := 0; i < v.n; i++ {
			bigger[i] = v.data[i]
			v.Copies++
		}
		v.data = bigger
	}
	v.data[v.n] = x
	v.n++
}

func (v *Vec) Get(i int) int { return v.data[i] }
func (v *Vec) Len() int      { return v.n }
func (v *Vec) Cap() int      { return len(v.data) }
