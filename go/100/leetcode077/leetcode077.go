package leetcode077

func combine(n int, k int) [][]int {
	var res [][]int
	helper(1, n, k, []int{}, &res)
	return res
}

func helper(start, n, k int, com []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, append([]int{}, com...))
		return
	}
	for i := start; i <= n; i++ {
		com = append(com, i)
		helper(i+1, n, k-1, com, res)
		com = com[:len(com)-1]
	}
}
