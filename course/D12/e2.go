package main

import "sort"

// Complete returns every inserted word that starts with prefix, in alphabetical order
// (an empty result when there is none). The prefix itself counts when it is a word.
//
// This is what an editor does when you type "co" and it offers const, continue, copy...
func Complete(t *Trie, prefix string) []string {
	start := t.find(prefix)
	if start == nil {
		return nil
	}
	var out []string
	// walk visits node, whose path from the root spells word
	var walk func(node *Trie, word string)
	walk = func(node *Trie, word string) {
		// hole 1: if a word ends at this node, add it to out

		// hole 2: visit the children in alphabetical order of their letter
		// (collect the letters, sort them, then call walk on each child with word+string(letter))
	}
	walk(start, prefix)
	_ = sort.Slice // Go refuses an unused import: delete this line once your code sorts the letters
	return out
}

// The trie from exercise 1, complete.

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
