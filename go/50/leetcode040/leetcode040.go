package leetcode040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return help(candidates, target)
}

func help(candidates []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] < target {
			for _, ele := range help(candidates[i+1:], target-candidates[i]) {
				tmp := make([]int, len(ele)+1)
				copy(tmp, ele)
				tmp[len(ele)] = candidates[i]
				res = append(res, tmp)
			}
		} else if candidates[i] == target {
			res = append(res, []int{candidates[i]})
		}
	}
	return res
}
