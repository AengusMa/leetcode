package leetcode027

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	valCount := 0
	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			valCount++
		} else {
			left++
		}
	}
	return len(nums) - valCount
}
