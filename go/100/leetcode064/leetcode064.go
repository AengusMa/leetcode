package leetcode064

// {
// {1, 3, 1},
// {1, 5, 1},
// {4, 2, 1},
// }
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
			// grid[i][j] += int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1])))
		}
	}
	return grid[n-1][m-1]
}
