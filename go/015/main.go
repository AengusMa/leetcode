package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		num := nums[i]
		if num > 0 {
			return result
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[j] + nums[k] + num
			if sum == 0 {
				result = append(result, []int{num, nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			}
		}
	}
	return result
}

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
