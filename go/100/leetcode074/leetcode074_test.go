package leetcode074

import "testing"

func Test074(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	if searchMatrix(matrix, 3) != true {
		t.Error("error")
	}
	if searchMatrix(matrix, 13) != false {
		t.Error("error")
	}
}
