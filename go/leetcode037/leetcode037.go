package leetcode037

func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	solve(board)
}

var bytes = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				for _, c := range bytes {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
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
			if v := board[i][j]; v != '.' {
				v -= '1'
				rows[i][v] = true
				cols[j][v] = true
				blocks[i/3][j/3][v] = true
			} else {
				positions = append(positions, [2]int{i, j})
			}
		}
	}
	dfs(board, rows, cols, blocks, positions, 0)
}

func dfs(board [][]byte, rows, cols [9][9]bool, blocks [3][3][9]bool, positions [][2]int, offset int) bool {
	if offset == len(positions) {
		return true
	}
	pos := positions[offset]
	row, col := pos[0], pos[1]
	blockRow := row / 3
	blockCol := col / 3
	for i := 0; i < 9; i++ {
		if !rows[row][i] && !cols[col][i] && !blocks[blockRow][blockCol][i] {
			// 尝试
			rows[row][i] = true
			cols[col][i] = true
			blocks[blockRow][blockCol][i] = true
			ok := dfs(board, rows, cols, blocks, positions, offset+1)
			if ok {
				board[row][col] = byte(i) + '1'
				return true
			}
			// 	回退
			rows[row][i] = false
			cols[col][i] = false
			blocks[blockRow][blockCol][i] = false
		}
	}
	return false
}
