package leetcode054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	top, left := 0, 0
	bottom, right := len(matrix)-1, len(matrix[0])-1
	ret := make([]int, 0)
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				ret = append(ret, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				ret = append(ret, matrix[i][left])
			}
			left++
		}
	}
	return ret
}
