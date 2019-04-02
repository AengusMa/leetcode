package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	result := 0
	sort.Ints(nums)
	min := 999999
	for i := 0; i < len(nums)-2; i++ {
		num := nums[i]
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[j] + nums[k] + num - target
			tmp := nums[j] + nums[k] + num
			if sum == 0 {
				return target
			} else if sum > 0 {
				k--
				if min > sum {
					min = sum
					result = tmp
				}
			} else if sum < 0 {
				j++
				if min > -sum {
					min = -sum
					result = tmp
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
func main() {
	nums := []int{-3, -2, -5, 3, -4}
	println(threeSumClosest(nums, -1))
	println(maxDepth(31))
}

// it returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
