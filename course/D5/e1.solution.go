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
	for _, e := range m.buckets[m.bucketOf(key)] {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

func (m *HashMap) Put(key string, value int) {
	i := m.bucketOf(key)
	for j := range m.buckets[i] {
		if m.buckets[i][j].key == key {
			m.buckets[i][j].value = value // already there: replace
			return
		}
	}
	m.buckets[i] = append(m.buckets[i], entry{key, value})
	m.count++
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
