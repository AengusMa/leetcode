package leetcode034

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			result = spiderResult(nums, target, mid)
			break
		}
	}
	return result
}

func spiderResult(nums []int, target, index int) []int {
	result := []int{index, index}
	tmp := index
	for tmp > 0 && nums[tmp-1] == target {
		tmp--
	}
	result[0] = tmp
	tmp = index
	for tmp < len(nums)-1 && nums[tmp+1] == target {
		tmp++
	}
	result[1] = tmp
	return result
}
