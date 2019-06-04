package leetcode119

func getRow(rowIndex int) []int {
	result := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		result = append(result, make([]int, i+1))
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
				continue
			}
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result[rowIndex]
}
