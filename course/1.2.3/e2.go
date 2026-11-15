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

// Browser keeps the pages behind and ahead of the current one.
type Browser struct {
	back, forward Stack
	Current       string
}

// Visit opens a new page: the current one goes behind, and nothing is ahead any more.
func (b *Browser) Visit(url string) {
	if b.Current != "" {
		b.back.Push(b.Current)
	}
	b.forward = Stack{}
	b.Current = url
}

// Back goes to the previous page. It returns false if there is none.
func (b *Browser) Back() bool {
	// hole
	return false
}

// Forward goes to the page Back left. It returns false if there is none.
func (b *Browser) Forward() bool {
	// hole
	return false
}
