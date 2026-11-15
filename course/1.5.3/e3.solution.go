package main

// LongestUnique returns the start and the length of the longest run of consecutive
// words with no word repeated. On a tie, the earliest run. An empty input gives 0, 0.
func LongestUnique(words []string) (start, length int) {
	last := map[string]int{} // word -> its last position
	left := 0
	for right, w := range words {
		if p, seen := last[w]; seen && p >= left {
			left = p + 1 // the repeat is inside the window: move past its first copy
		}
		last[w] = right
		if right-left+1 > length {
			start, length = left, right-left+1
		}
	}
	return start, length
}
