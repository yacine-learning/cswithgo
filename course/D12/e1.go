package main

// Trie stores words letter by letter: every path down from the root spells a prefix.
type Trie struct {
	children map[rune]*Trie
	word     bool // a word ends at this node
}

func NewTrie() *Trie { return &Trie{children: map[rune]*Trie{}} }

// Insert adds w to the trie.
func (t *Trie) Insert(w string) {
	node := t
	for _, r := range w {
		// Go refuses to compile a variable that is never used. These lines "use" r and node
		// by throwing the values away, so the file compiles before you fill the holes.
		// Delete them once your code uses r and node.
		_, _ = r, node

		// hole 1: go down to the child for r, creating it when it is missing
	}
	// hole 2: mark that a word ends here
}

// Contains reports whether w itself was inserted.
func (t *Trie) Contains(w string) bool {
	n := t.find(w)
	return n != nil && n.word
}

// HasPrefix reports whether some inserted word starts with p.
func (t *Trie) HasPrefix(p string) bool { return t.find(p) != nil }

// find follows the letters of s from the root. It returns the node where s ends,
// or nil when the path stops before the end of s.
func (t *Trie) find(s string) *Trie {
	node := t
	for _, r := range s {
		_ = r // delete this line once your code uses r (see above)

		// hole 3: step to the child for r; if there is none, s is not in the trie
	}
	return node
}
