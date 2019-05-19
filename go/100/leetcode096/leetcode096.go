package leetcode096

func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			f[i] += f[j] * f[i-1-j]
		}
	}
	return f[n]
}
