package leetcode051

func solveNQueens(n int) [][]string {
	// init board
	board := BoardConstructor(n)
	// result
	var result [][]string

	var backtracking func(b *Board, col, n int)
	backtracking = func(b *Board, col, n int) {
		if col == n {
			// if the col reaches to n, it means we've just found a result
			result = append(result, b.Clone())
			return
		}
		// try every row in the next column
		for i := 0; i < n; i++ {
			if b.IsSafe(i, col) {
				// place a queen
				b.Place(i, col)
				// try next column
				backtracking(b, col+1, n)
				// remove a queen
				b.Remove(i, col)
			}
		}
	}

	backtracking(&board, 0, n)
	return result
}

// Board board class
type Board struct {
	M []string
	N int
}

func BoardConstructor(n int) Board {
	var m []string
	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			temp += "."
		}
		m = append(m, temp)
	}
	return Board{m, n}
}

func (b *Board) Place(row, col int) {
	// it means b.M[row][col] = "Q"
	b.M[row] = b.M[row][:col] + "Q" + b.M[row][col+1:]
}

func (b *Board) Remove(row, col int) {
	// it means b.M[row][col] = "."
	b.M[row] = b.M[row][:col] + "." + b.M[row][col+1:]
}

func (b *Board) IsSafe(row, col int) bool {
	// check row and col
	for i := 0; i < b.N; i++ {
		if b.M[i][col] == 'Q' {
			return false
		}
		if b.M[row][i] == 'Q' {
			return false
		}
	}
	// check diagonal
	for i := 0; i < b.N; i++ {
		for j := 0; j < b.N; j++ {
			if i+j == row+col || i-j == row-col {
				if i != row && j != col && b.M[i][j] == 'Q' {
					return false
				}
			}
		}
	}
	return true
}

func (b *Board) Clone() []string {
	var m []string
	for i := 0; i < b.N; i++ {
		m = append(m, b.M[i])
	}
	return m
}
