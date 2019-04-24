package leetcode053

func maxSubArray(nums []int) int {
	max := nums[0]
	keepMax := nums[0]
	for i := 1; i < len(nums); i++ {
		keepMax = maxNum(nums[i]+keepMax, nums[i])
		max = maxNum(keepMax, max)
	}
	return max
}
func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
