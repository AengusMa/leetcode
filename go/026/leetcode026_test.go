package leetcode026

import (
	"fmt"
	"testing"
)

func Test026(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	for _, ele := range nums {
		println(ele)
	}
}
