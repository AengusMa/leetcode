package leetcode036

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidSquare(board, i, i+1, 0, 9) {
			return false
		}
		if !isValidSquare(board, 0, 9, i, i+1) {
			return false
		}
	}
	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			if !isValidSquare(board, i, i+3, j, j+3) {
				return false
			}
		}
	}
	return true
}
func isValidSquare(board [][]byte, rowBegin, row, colBegin, col int) bool {
	tmp := make([]bool, 9)
	for i := rowBegin; i < row; i++ {
		for j := colBegin; j < col; j++ {
			if board[i][j] != '.' {
				if tmp[board[i][j]-49] {
					return false
				}
				tmp[board[i][j]-49] = true
			}
		}
	}
	return true
}
