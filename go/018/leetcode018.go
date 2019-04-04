package leetcode018

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums)-1; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for n := k + 1; n < len(nums); n++ {
					if n > k+1 && nums[n] == nums[n-1] {
						continue
					}
					if nums[i]+nums[j]+nums[k]+nums[n] == target {
						tmp := []int{nums[i], nums[j], nums[k], nums[n]}
						result = append(result, tmp)
					}
				}
			}
		}
	}
	return result
}

func fourSum1(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, len(nums)-1; k < l; {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return result
}
