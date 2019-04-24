package leetcode045

// https://leetcode.windliang.cc/leetCode-45-Jump-Game-II.html
// 贪心
func jump(nums []int) int {
	ret, last, curr := 0, 0, 0
	for index, num := range nums {
		if index > last {
			last = curr
			ret++
		}
		curr = max(curr, num+index)
	}
	return ret
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
