package leetcode078

func subsets(nums []int) [][]int {
	var res [][]int
	help(nums, []int{}, 0, &res)
	return res
}

func help(nums, tmp []int, start int, res *[][]int) {
	*res = append(*res, append([]int{}, tmp...))
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help(nums, tmp, i+1, res)
		tmp = tmp[:len(tmp)-1]
	}
}
