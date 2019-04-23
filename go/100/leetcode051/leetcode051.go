package leetcode051

func solveNQueens(n int) [][]string {
	ret := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp []string
		for i := 0; i < n; i++ {
			tmp = append(tmp, ".")
		}
		ret[i] = tmp
	}

	return ret
}
