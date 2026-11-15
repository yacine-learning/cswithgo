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
		next, ok := node.children[r]
		if !ok {
			next = NewTrie()
			node.children[r] = next
		}
		node = next
	}
	node.word = true
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
		next, ok := node.children[r]
		if !ok {
			return nil
		}
		node = next
	}
	return node
}
