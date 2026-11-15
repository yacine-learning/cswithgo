package main

import "fmt"

func main() {
	l := FromValues(7, 7, 1, 2, 7, 3, 7)
	fmt.Println("list:             ", l.Values())
	l.RemoveAll(7)
	fmt.Println("RemoveAll(7):     ", l.Values(), " (first, middle and last removed alike)")
	l.PushFront(9)
	fmt.Println("PushFront(9):     ", l.Values())
	l.RemoveAll(3)
	fmt.Println("RemoveAll(3):     ", l.Values())
	e := FromValues(5, 5, 5)
	e.RemoveAll(5)
	fmt.Println("all removed:      ", e.Values(), " (the dummy head is still there)")
}
