package leetcode050

import (
	"fmt"
	"testing"
)

func Test050(t *testing.T) {
	fmt.Println(myPow(2.00000, 10))
	if myPow(2.00000, 10) != 1024 {
		t.Error("出错了")
	}
	if myPow(2.00000, -2) != 0.25 {
		t.Error("出错了")
	}
}
