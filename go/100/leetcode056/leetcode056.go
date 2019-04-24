package leetcode056

import "sort"

//  [[1,3],[2,6],[8,10],[15,18]]
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		num1 := intervals[i]
		num2 := intervals[j]
		if num1[0] == num2[0] {
			return num1[1] < num2[1]
		}
		return num1[0] < num2[0]
	})
	var ret [][]int
	current := intervals[0]
	for _, next := range intervals[1:] {
		if current[1] < next[0] {
			ret = append(ret, current)
			current = next
		} else {
			current[1] = max(next[1], current[1])
		}
	}
	ret = append(ret, current)
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
