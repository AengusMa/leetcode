package leetcode062

func uniquePaths(m int, n int) int {
	var dp [][]int
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			tmp[j] = 1
		}
		dp = append(dp, tmp)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	// fmt.Println(dp)
	return dp[n-1][m-1]
}

func help(i, j, m, n int) int {
	if i == m && j == n {
		return 1
	}
	if i > m || j > n {
		return 0
	}
	ret := 0
	if i != m {
		i++
		ret += help(i, j, m, n)
		i--
	}
	if j != n {
		j++
		ret += help(i, j, m, n)
		j--
	}
	return ret
}
