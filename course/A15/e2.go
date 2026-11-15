package main

// Suggest returns the known word closest to word, for a "did you mean …?" message.
// Only words at distance 1 or 2 are worth suggesting; on a tie, the first in alphabetical
// order wins. ok is false when no known word is close enough (or word is itself known).
func Suggest(word string, known []string) (best string, ok bool) {
	bestDist := 3 // anything farther than 2 is not a suggestion
	for _, k := range known {
		_ = k // Go refuses unused variables: delete this line once your code uses k

		// hole: if word is known, there is nothing to suggest.
		//       Otherwise keep k when it is closer than bestDist, or as close but
		//       alphabetically first.
	}
	_ = bestDist // delete this line once your code uses bestDist
	return best, ok
}

// EditDistance from exercise 1, complete.

// EditDistance is the fewest single-letter changes (insert, delete or replace a letter) that
// turn a into b: Levenshtein, 1965.
//
// d[i][j] is the distance between the first i letters of a and the first j letters of b.
// Each cell only needs three cells already filled: the one above, the one to the left,
// and the one diagonally above-left. Fill the table row by row; the answer is the last cell.
func EditDistance(a, b string) int {
	x, y := []rune(a), []rune(b)
	d := make([][]int, len(x)+1)
	for i := range d {
		d[i] = make([]int, len(y)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			change := 1
			if x[i-1] == y[j-1] {
				change = 0
			}
			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+change)
		}
	}
	return d[len(x)][len(y)]
}
