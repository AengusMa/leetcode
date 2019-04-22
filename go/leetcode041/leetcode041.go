package leetcode041

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	res := 0
	for ; res < len(nums) && nums[res] == res+1; res++ {
	}
	return res + 1
}

// 标记法
func firstMissingPositive1(nums []int) int {
	// 将p之前的都交换为正数
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	// 对应位置有值变成负数
	for i := 0; i < p; i++ {
		index := abs(nums[i]) - 1
		if index < p {
			nums[index] = -abs(nums[index])
		}
	}
	for i := 0; i < p; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return p + 1
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
