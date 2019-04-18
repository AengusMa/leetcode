package leetcode032

import (
	"testing"
)

func Test032(t *testing.T) {
	if longestValidParentheses("())())()()()()") != 8 {
		t.Error("error")
	}
	if longestValidParentheses("()(())") != 6 {
		t.Error("error")
	}
	if longestValidParentheses(")()())") != 4 {
		t.Error("error")
	}
	if longestValidParenthesesDP("())())()()()()") != 8 {
		t.Error("error")
	}
	if longestValidParenthesesDP("()(())") != 6 {
		t.Error("error")
	}
	if longestValidParenthesesDP(")()())") != 4 {
		t.Error("error")
	}
}
