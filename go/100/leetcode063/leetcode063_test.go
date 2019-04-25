package leetcode063

import (
	"fmt"
	"testing"
)

func Test063(t *testing.T) {
	nums := [][]int{
		{0},
		{0},
	}
	fmt.Println(uniquePathsWithObstacles(nums))
	nums = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	if uniquePathsWithObstacles(nums) != 2 {
		t.Error("error")
	}
}
