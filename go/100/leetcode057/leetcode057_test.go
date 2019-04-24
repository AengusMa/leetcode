package leetcode057

import (
	"fmt"
	"testing"
)

func Test057(t *testing.T) {
	nums := [][]int{
		{1, 3},
		{6, 9},
	}
	fmt.Println(insert(nums, []int{2, 5}))
	// [[1,2],[3,5],[6,7],[8,10],[12,16]]
	//            [4,8]
	// [[1,2],[3,10],[12,16]]
	nums = [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	fmt.Println(insert(nums, []int{4, 8}))
}
