package main

type entry struct {
	key   string
	value int
}

// HashMap maps strings to ints with buckets of entries ("chaining").
type HashMap struct {
	buckets [][]entry
	count   int
}

func NewHashMap() *HashMap { return &HashMap{buckets: make([][]entry, 4)} }

// hash is FNV-1a: mix each byte into the number. Same key, same number, always.
func hash(key string) uint32 {
	h := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return h
}

func (m *HashMap) bucketOf(key string) int { return int(hash(key) % uint32(len(m.buckets))) }

func (m *HashMap) Get(key string) (int, bool) {
	// hole: look through the entries of key's bucket; return the value of the one with this key
	return 0, false
}

func (m *HashMap) Put(key string, value int) {
	i := m.bucketOf(key)
	// hole 1: if an entry of bucket i has this key, replace its value and return

	// hole 2: otherwise add entry{key, value} to bucket i and count it
	_ = i // delete this line once you use i

	if m.count > 2*len(m.buckets) { // more than 2 entries per bucket on average: grow
		m.grow()
	}
}

// grow doubles the buckets and puts every entry in its new bucket.
func (m *HashMap) grow() {
	old := m.buckets
	m.buckets = make([][]entry, 2*len(old))
	for _, b := range old {
		for _, e := range b {
			i := m.bucketOf(e.key)
			m.buckets[i] = append(m.buckets[i], e)
		}
	}
}

func (m *HashMap) Len() int { return m.count }
