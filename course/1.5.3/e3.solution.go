package main

// LongestUnique returns the start and the length of the longest run of consecutive
// IDs with no ID repeated. On a tie, the earliest run. An empty input gives 0, 0.
func LongestUnique(ids []string) (start, length int) {
	last := map[string]int{} // ID -> its last position
	left := 0
	for right, id := range ids {
		if p, seen := last[id]; seen && p >= left {
			left = p + 1 // the repeat is inside the window: move past its first copy
		}
		last[id] = right
		if right-left+1 > length {
			start, length = left, right-left+1
		}
	}
	return start, length
}
