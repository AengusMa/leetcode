package leetcode040

import (
	"fmt"
	"testing"
)

func Test040(t *testing.T) {
	// [
	// 	[1, 7],
	// 	[1, 2, 5],
	// 	[2, 6],
	// 	[1, 1, 6]
	// ]
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// [
	// 	[1,2,2],
	// 	[5]
	// ]
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
