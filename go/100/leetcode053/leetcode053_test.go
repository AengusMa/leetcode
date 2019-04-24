package leetcode053

import (
	"testing"
)

func Test053(t *testing.T) {
	// [4,-1,2,1] has the largest sum = 6
	if maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("出错了")
	}
}
