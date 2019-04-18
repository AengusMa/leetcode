package leetcode031

func nextPermutation(nums []int) {
	tmp := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			tmp = i
			break
		}
	}
	if tmp == -1 {
		reverseSort(nums)
		return
	}
	for j := len(nums) - 1; j > tmp; j-- {
		if nums[j] > nums[tmp] {
			nums[j], nums[tmp] = nums[tmp], nums[j]
			reverseSort(nums[tmp+1:])
			break
		}
	}
}
func reverseSort(nums []int) {
	start := 0
	end := len(nums) - 1
	for end > start {
		nums[end], nums[start] = nums[start], nums[end]
		end--
		start++
	}
}
