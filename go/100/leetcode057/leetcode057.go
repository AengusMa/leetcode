package leetcode057

// [[1,2],[3,5],[6,7],[8,10],[12,16]]
//            [3,8]
// [[1,2],[3,10],[12,16]]
func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int
	var tmp []int
	for i := 0; i < len(intervals)-1; i++ {
		if tmp != nil {
			if tmp[1] >= intervals[i][0] {
				tmp[1] = max(tmp[1], intervals[i][1])
			} else {
				ret = append(ret, tmp)
				tmp = nil
			}
		}
		if newInterval[0] < intervals[i+1][0] {
			tmp = make([]int, 2)
			tmp[0] = newInterval[0]
			tmp[1] = newInterval[1]
			if newInterval[0] <= intervals[i][1] {
				tmp[0] = intervals[i][0]
				tmp[1] = max(intervals[i][1], tmp[1])
			}
		} else {
			ret = append(ret, intervals[i])
		}
	}
	if tmp != nil {
		ret = append(ret, tmp)
	}
	return ret
}

// {1, 3},
// {6, 9},
// 2  5
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
