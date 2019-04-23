package leetcode037

func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	solve1(board)
}

var bytes = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solve1(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				for _, c := range bytes {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve1(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
func isValid(board [][]byte, row, col int, b byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == b || board[row][i] == b {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == b {
			return false
		}
	}
	return true
}

func solveSudoku1(board [][]byte) {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	blocks := [3][3][9]bool{}

	var positions [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ele := board[i][j]; ele != '.' {
				ele -= '1'
				rows[i][ele] = true
				cols[j][ele] = true
				blocks[i/3][j/3][ele] = true
			} else {
				positions = append(positions, [2]int{i, j})
			}
		}
	}
	solve(board, rows, cols, blocks, positions, 0)

}

func solve(board [][]byte, rows, cols [9][9]bool, blocks [3][3][9]bool, positions [][2]int, offset int) bool {
	if len(positions) == offset {
		return true
	}
	row := positions[offset][0]
	col := positions[offset][1]
	for i := 0; i < 9; i++ {
		if !rows[row][i] && !cols[col][i] && !blocks[row/3][col/3][i] {
			rows[row][i] = true
			cols[col][i] = true
			blocks[row/3][col/3][i] = true
			ok := solve(board, rows, cols, blocks, positions, offset+1)
			if ok {
				board[row][col] = byte(i) + '1'
				return true
			}
			rows[row][i] = false
			cols[col][i] = false
			blocks[row/3][col/3][i] = false
		}
	}
	return false
}
