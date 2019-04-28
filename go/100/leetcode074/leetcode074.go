package leetcode074

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1

	for left <= right {
		if matrix[left/n][left%n] == target || matrix[right/n][right%n] == target {
			return true
		}
		mid := (left + right) >> 1
		i := mid / n
		j := mid % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
