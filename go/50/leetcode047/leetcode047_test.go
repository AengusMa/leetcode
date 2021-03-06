package leetcode047

import (
	"fmt"
	"testing"
)

func Test047(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
	// [[0,0,0,1,9],[0,0,0,9,1],[0,0,1,0,9],[0,0,1,9,0],
	// [0,0,9,0,1],[0,0,9,1,0],[0,1,0,0,9],[0,1,0,9,0],
	// [0,1,9,0,0],[0,9,0,0,1],[0,9,0,1,0],[0,9,1,0,0],
	// [1,0,0,0,9],[1,0,0,9,0],[1,0,9,0,0],[1,9,0,0,0],
	// [9,0,0,0,1],[9,0,0,1,0],[9,0,1,0,0],[9,1,0,0,0]]
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}
