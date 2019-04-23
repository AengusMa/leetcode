package leetcode041

import (
	"testing"
)

func Test041(t *testing.T) {

	if firstMissingPositive1([]int{1, 1, 2, 2, 3, 4, 2}) != 5 {
		t.Error("出错了")
	}
	if firstMissingPositive1([]int{1, 2, 0}) != 3 {
		t.Error("出错了")
	}
	if firstMissingPositive1([]int{3, 4, -1, 1}) != 2 {
		t.Error("出错了")
	}
	if firstMissingPositive1([]int{12, 11, 9, 7, 8}) != 1 {
		t.Error("出错了")
	}
}
