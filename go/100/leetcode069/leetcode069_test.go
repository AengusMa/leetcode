package leetcode069

import (
	"fmt"
	"testing"
)

func Test069(t *testing.T) {
	if mySqrt(4) != 2 {
		t.Error("error")
	}
	if mySqrt(8) != 2 {
		t.Error("error")
	}
	fmt.Println(mySqrt(1212195572))
}
