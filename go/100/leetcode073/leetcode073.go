package leetcode073

func setZeroes(matrix [][]int) {
	row := make(map[int]struct{}, 0)
	col := make(map[int]struct{}, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}
	for key := range row {
		for i := 0; i < len(matrix[key]); i++ {
			matrix[key][i] = 0
		}
	}
	for key := range col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][key] = 0
		}
	}
}
