package leetcode032

import (
	"container/list"
)

func longestValidParentheses(s string) int {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		ele := stack.Back()
		if ele == nil {
			stack.PushBack(i)
		} else {
			if s[ele.Value.(int)] == '(' && s[i] == ')' {
				stack.Remove(ele)
			} else {
				stack.PushBack(i)
			}
		}
	}
	if stack.Len() == 0 {
		return len(s)
	}
	result := 0
	last := len(s)
	for stack.Len() != 0 {
		ele := stack.Back()
		result = max(result, last-ele.Value.(int)-1)
		last = ele.Value.(int)
		stack.Remove(ele)
	}
	return max(result, last)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ( ) ( ) ( ( ) )
// 0 1 2 3 4 5 6 7
func longestValidParenthesesDP(s string) int {
	if len(s) == 1 {
		return 0
	}
	result := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = 2 + dp[i-2]
				}
			}
			if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				if i-2-dp[i-1] >= 0 {
					dp[i] = 2 + dp[i-1] + dp[i-2-dp[i-1]]
				} else {
					dp[i] = 2 + dp[i-1]
				}
			}
			result = max(result, dp[i])
		}
	}
	return result
}
