package leetcode035

func searchInsert1(nums []int, target int) int {
	for index, num := range nums {
		if target <= num {
			return index
		}
	}
	return len(nums)
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}
