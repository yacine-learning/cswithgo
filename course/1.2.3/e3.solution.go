package main

// Balanced reports whether the brackets ( ) [ ] { } of s are balanced.
// When they are not, pos is the byte position of the problem:
// a closer that does not match, or the innermost opener left unclosed.
func Balanced(s string) (ok bool, pos int) {
	opener := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []int // positions of the openers not closed yet
	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '(', '[', '{':
			stack = append(stack, i)
		case ')', ']', '}':
			if len(stack) == 0 || s[stack[len(stack)-1]] != opener[c] {
				return false, i
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false, stack[len(stack)-1]
	}
	return true, -1
}
