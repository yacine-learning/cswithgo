package main

import "fmt"

func num(v int) *Node             { return &Node{Num: v} }
func op(o rune, l, r *Node) *Node { return &Node{Op: o, Left: l, Right: r} }

func main() {
	e := op('*', num(3), op('+', num(4), op('-', num(5), num(1))))
	try("a single number", func() string {
		if Eval(num(7)) != 7 || Infix(num(7)) != "7" {
			return fmt.Sprintf("Eval = %d, Infix = %q for the number 7", Eval(num(7)), Infix(num(7)))
		}
		return ""
	})
	try("Eval 3 * (4 + (5 - 1))", func() string {
		if v := Eval(e); v != 24 {
			return fmt.Sprintf("Eval = %d, expected 24", v)
		}
		return ""
	})
	try("Infix 3 * (4 + (5 - 1))", func() string {
		if s := Infix(e); s != "(3 * (4 + (5 - 1)))" {
			return fmt.Sprintf("Infix = %q, expected %q", s, "(3 * (4 + (5 - 1)))")
		}
		return ""
	})
	try("order matters for - and /", func() string {
		x := op('/', op('-', num(20), num(8)), num(4))
		if v := Eval(x); v != 3 {
			return fmt.Sprintf("Eval((20 - 8) / 4) = %d, expected 3: left operand first", v)
		}
		return ""
	})
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
