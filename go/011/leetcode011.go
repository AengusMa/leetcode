package leetcode011

func maxArea(height []int) int {
	result := 0
	r := len(height) - 1
	l := 0
	for l < r {
		rh := height[r]
		lh := height[l]
		w := r - l
		result = max(result, min(rh, lh)*w)
		if rh < lh {
			r--
		} else {
			l++
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
