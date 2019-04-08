package leetcode027

import (
	"fmt"
	"testing"
)

func Test027(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 2))
	for _, ele := range nums {
		println(ele)
	}
}
