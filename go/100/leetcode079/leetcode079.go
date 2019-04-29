package leetcode079

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if help(board, []byte(word), i, j) {
					return true
				}
			}
		}
	}
	return false
}

func help(board [][]byte, word []byte, m, n int) bool {
	if len(word) == 0 {
		return true
	}
	if m >= len(board) || n >= len(board[0]) || n < 0 || m < 0 {
		return false
	}
	if board[m][n] != word[0] {
		return false
	}
	board[m][n] = board[m][n] ^ 255
	exist := help(board, word[1:], m+1, n) || help(board, word[1:], m-1, n) || help(board, word[1:], m, n+1) || help(board, word[1:], m, n-1)
	board[m][n] = board[m][n] ^ 255
	return exist
}
