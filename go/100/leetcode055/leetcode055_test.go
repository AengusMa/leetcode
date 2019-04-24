package leetcode055

import (
	"fmt"
	"testing"
)

func Test055(t *testing.T) {
	if canJump([]int{2, 3, 1, 1, 4}) != true {
		t.Error("出错了")
	}
	if canJump([]int{3, 2, 1, 0, 4}) != false {
		t.Error("出错了")
	}
	if canJump([]int{1, 2, 0, 1}) != true {
		t.Error("出错了")
	}
	fmt.Println(canJump([]int{4, 6, 13, 4, 6, 9}))
}
