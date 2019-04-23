package leetcode039

import (
	"fmt"
	"testing"
)

func Test039(t *testing.T) {
	// [
	// 	[7],
	// 	[2,2,3]
	// ]
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	// [
	// 	[2,2,2,2],
	// 	[2,3,3],
	// 	[3,5]
	// ]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
