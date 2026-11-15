package main

// Stack is a stack of strings on a slice: the top is the last element.
type Stack struct {
	items []string
}

func (s *Stack) Push(x string) { s.items = append(s.items, x) }

// Pop removes and returns the top. ok is false when the stack is empty.
func (s *Stack) Pop() (x string, ok bool) {
	if len(s.items) == 0 {
		return "", false
	}
	x = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return x, true
}

// Peek returns the top without removing it.
func (s *Stack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) Len() int { return len(s.items) }
