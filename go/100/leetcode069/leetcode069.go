package leetcode069

import "sort"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left := 1
	right := x - 1
	for left < right-1 {
		mid := (left + right) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}
func mySqrt1(x int) int {
	return sort.Search(x, func(i int) bool { return (i+1)*(i+1) > x })
}
