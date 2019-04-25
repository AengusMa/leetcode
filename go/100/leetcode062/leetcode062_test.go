package leetcode062

import (
	"fmt"
	"testing"
)

func Test062(t *testing.T) {
	fmt.Println(uniquePaths(3, 2))
	if uniquePaths(3, 2) != 3 {
		t.Error("error")
	}
	fmt.Println(uniquePaths(7, 3))
	if uniquePaths(7, 3) != 28 {
		t.Error("error")
	}
}
