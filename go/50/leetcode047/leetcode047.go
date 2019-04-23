package leetcode047

func permuteUnique(nums []int) [][]int {
	return helper(nums, 0)
}

func helper(nums []int, cur int) [][]int {
	if cur == len(nums) {
		p := append([]int{}, nums...)
		return [][]int{p}
	}
	ret := make([][]int, 0)
	visited := make(map[int]struct{})
	for i := cur; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; !ok {
			visited[nums[i]] = struct{}{}
			nums[cur], nums[i] = nums[i], nums[cur]
			ret = append(ret, helper(nums, cur+1)...)
			nums[cur], nums[i] = nums[i], nums[cur]
		}

	}
	return ret
}
