package leetcode046

func permute1(nums []int) [][]int {
	var ret [][]int
	var helper func([]int)
	helper = func(tmp []int) {
		if len(tmp) == len(nums) {
			tm := append([]int{}, nums...)
			ret = append(ret, tm)
			return
		}
		for i := 0; i < len(nums); i++ {
			if contains(tmp, nums[i]) {
				continue
			}
			tmp = append(tmp, nums[i])
			helper(tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	helper([]int{})
	return ret
}
func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// 基于上面的改良
func permute(nums []int) [][]int {
	var ret [][]int
	var helper func(int)
	helper = func(cur int) {
		if cur == len(nums) {
			tmp := append([]int{}, nums...)
			ret = append(ret, tmp)
			return
		}
		for i := cur; i < len(nums); i++ {
			// 将临时的结果放在cur前面
			nums[i], nums[cur] = nums[cur], nums[i]
			helper(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}
	helper(0)
	return ret
}
func permute2(nums []int) [][]int {
	return help(nums, 0)
}
func help(nums []int, cur int) [][]int {
	if cur == len(nums) {
		p := append([]int{}, nums...)
		return [][]int{p}
	}
	res := make([][]int, 0)
	for i := cur; i < len(nums); i++ {
		nums[cur], nums[i] = nums[i], nums[cur]
		res = append(res, help(nums, cur+1)...)
		nums[cur], nums[i] = nums[i], nums[cur]
	}
	return res
}
