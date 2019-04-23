package leetcode034

import (
	"fmt"
	"testing"
)

func Test034(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 5))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{2, 2}, 2))
}
