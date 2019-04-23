package leetcode048

func rotate(matrix [][]int) {
	// reverse
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
