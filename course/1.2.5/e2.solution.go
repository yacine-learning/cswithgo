package main

type node struct {
	value int
	next  *node
}

// List is a singly linked list. head is a dummy node: the first value is head.next.
type List struct {
	head node
}

// PushFront adds v before the first value: O(1).
func (l *List) PushFront(v int) {
	l.head.next = &node{value: v, next: l.head.next}
}

// RemoveAll removes every node holding v. Thanks to the dummy head,
// removing the first value is the same as removing any other.
func (l *List) RemoveAll(v int) {
	prev := &l.head
	for prev.next != nil {
		if prev.next.value == v {
			prev.next = prev.next.next // skip it; do not move prev: the new next may match too
		} else {
			prev = prev.next
		}
	}
}

// Values returns the values from first to last: O(n).
func (l *List) Values() []int {
	var out []int
	for n := l.head.next; n != nil; n = n.next {
		out = append(out, n.value)
	}
	return out
}

// FromValues builds a list holding vs in order.
func FromValues(vs ...int) *List {
	l := &List{}
	for i := len(vs) - 1; i >= 0; i-- {
		l.PushFront(vs[i])
	}
	return l
}

// Reverse reverses the list in place, without creating nodes.
func (l *List) Reverse() {
	var prev *node
	cur := l.head.next
	for cur != nil {
		next := cur.next // save it before changing cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head.next = prev
}
