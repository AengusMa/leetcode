package leetcode118

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		numberOfItems := i + 1
		ret[i] = make([]int, numberOfItems)
		for j := 0; j < numberOfItems; j++ {
			if j == 0 || j == numberOfItems-1 {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
	}
	return ret
}
