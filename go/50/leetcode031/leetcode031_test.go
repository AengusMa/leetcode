package leetcode031

import (
	"fmt"
	"testing"
)

func Test031(t *testing.T) {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
