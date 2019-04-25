package leetcode070

import (
	"testing"
)

func Test070(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Error("error")
	}
	if climbStairs(3) != 3 {
		t.Error("error")
	}
	if climbStairs(44) != 1134903170 {
		t.Error("error")
	}
}
