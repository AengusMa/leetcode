package leetcode039

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(candidates); i++ {
		if candidates[i] < target {
			for _, ele := range combinationSum(candidates[i:], target-candidates[i]) {
				tmp := make([]int, len(ele)+1)
				copy(tmp, ele)
				tmp[len(ele)] = candidates[i]
				result = append(result, tmp)
			}
		} else if candidates[i] == target {
			result = append(result, []int{candidates[i]})
		}
	}
	return result
}
