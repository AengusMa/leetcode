package leetcode042

func trap(height []int) int {
	sum, maxLeft, maxRight := 0, 0, 0
	left, right := 0, len(height)-1
	for left <= right {
		if height[left] <= height[right] {
			if maxLeft <= height[left] {
				maxLeft = height[left]
			} else {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			if maxRight <= height[right] {
				maxRight = height[right]
			} else {
				sum += maxRight - height[right]
			}
			right--
		}
	}
	return sum
}
