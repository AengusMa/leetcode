package leetcode042

import "testing"

func Test042(t *testing.T) {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	if trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) != 6 {
		t.Error("错误")
	}
	if trap([]int{0}) != 0 {
		t.Error("错误")
	}
	if trap([]int{1, 7, 8}) != 0 {
		t.Error("错误")
	}
	// 5   2   5
	if trap([]int{5, 2, 1, 2, 1, 5}) != 14 {
		t.Error("错误")
	}
}
