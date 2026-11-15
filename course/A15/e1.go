package main

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
	// hole 1: the first column and the first row. Turning i letters into nothing takes
	//         i deletions; turning nothing into j letters takes j insertions.

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			// hole 2: the cheapest of
			//   d[i-1][j] + 1      delete x[i-1]
			//   d[i][j-1] + 1      insert y[j-1]
			//   d[i-1][j-1] + 0 or 1   keep the letter if x[i-1] == y[j-1], replace it otherwise
		}
	}
	return d[len(x)][len(y)]
}
