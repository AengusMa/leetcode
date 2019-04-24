package leetcode058

import (
	"testing"
)

func Test058(t *testing.T) {
	if lengthOfLastWord("Hello World") != 5 {
		t.Error("error")
	}
	if lengthOfLastWord("a  ") != 1 {
		t.Error("error")
	}
}
