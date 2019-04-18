package leetcode033

import "testing"

func Test033(t *testing.T) {
	if search([]int{4, 5, 6, 7, 0, 1, 2}, 0) != 4 {
		t.Error("error")
	}
	if search([]int{4, 5, 6, 7, 0, 1, 2}, 3) != -1 {
		t.Error("error")
	}
}
