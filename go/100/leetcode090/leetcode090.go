package leetcode090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	help(nums, []int{}, 0, &res)
	return res
}

func help(nums, tmp []int, start int, res *[][]int) {
	if start <= len(nums) {
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := start; i < len(nums); {
		tmp = append(tmp, nums[i])
		help(nums, tmp, i+1, res)
		tmp = tmp[:len(tmp)-1]
		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
}
