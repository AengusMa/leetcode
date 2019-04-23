package leetcode033

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, target, 0, len(nums)-1)
}

//4, 5, 6, 7, 10, 1, 2
//6, 7, 0, 1, 2, 3, 4, 5
func binarySearch(nums []int, target, left, right int) int {
	if nums[left] == target {
		return left
	}
	if left >= right {
		return -1
	}
	if nums[right] == target {
		return right
	}
	if right == left+1 {
		return -1
	}

	mid := (left + right) >> 1
	if nums[mid] == target {
		return mid
	}
	//target在有序的列表中
	if nums[left] < target && nums[mid] > target {
		return binarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < target && nums[right] > target {
		return binarySearch(nums, target, mid+1, right)
	}
	//target在无序的列表中
	if nums[mid] > nums[right] {
		return binarySearch(nums, target, mid+1, right)
	}
	return binarySearch(nums, target, left, mid-1)
}
