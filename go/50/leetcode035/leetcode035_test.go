package leetcode035

import (
	"fmt"
	"testing"
)

func Test035(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 7}, 7))
	if searchInsert([]int{1, 3, 5, 6}, 5) != 2 {
		t.Error("error")
	}
	if searchInsert([]int{1, 3, 5, 6}, 2) != 1 {
		t.Error("error")
	}
	if searchInsert([]int{1, 3, 5, 6}, 7) != 4 {
		t.Error("error")
	}
	if searchInsert([]int{1, 3, 5, 6}, 0) != 0 {
		t.Error("error")
	}
}
