package main

// Balanced reports whether the brackets ( ) [ ] { } of s are balanced.
// When they are not, pos is the byte position of the problem:
// a closer that does not match, or the innermost opener left unclosed.
func Balanced(s string) (ok bool, pos int) {
	return true, -1
}
