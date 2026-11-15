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
		// hole 1: if r == n, every row has a queen: one more solution
		for c := 0; c < n; c++ {
			// hole 2: skip c if its column or one of its diagonals is taken
			// hole 3: mark the three, place the rest (row r+1), then unmark the three
		}
	}
	place(0)
	_, _, _ = col, diag, anti // Go refuses unused variables: delete once your code uses them
	return count
}
