package leetcode056

import (
	"fmt"
	"testing"
)

func Test056(t *testing.T) {
	nums := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	// [[1,6],[8,10],[15,18]]
	fmt.Println(merge(nums))
}
