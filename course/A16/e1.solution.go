package main

// Queens counts the ways to place n queens on an n×n board so that no two attack each other
// (same column, or same diagonal; one queen per row by construction).
//
// Backtracking: place a queen in row r, in each free column in turn; go on to row r+1; when
// that returns, take the queen back and try the next column. A dead end costs one step back,
// not a restart.
func Queens(n int) int {
	col := make([]bool, n)    // col[c]: a queen is in column c
	diag := make([]bool, 2*n) // diag[r+c]: a queen is on this ↗ diagonal
	anti := make([]bool, 2*n) // anti[r-c+n]: a queen is on this ↘ diagonal
	count := 0
	var place func(r int)
	place = func(r int) {
		if r == n {
			count++
			return
		}
		for c := 0; c < n; c++ {
			if col[c] || diag[r+c] || anti[r-c+n] {
				continue
			}
			col[c], diag[r+c], anti[r-c+n] = true, true, true
			place(r + 1)
			col[c], diag[r+c], anti[r-c+n] = false, false, false
		}
	}
	place(0)
	return count
}
