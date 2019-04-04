package leetcode020

import (
	"container/list"
)

func isValid(s string) bool {
	st := list.New()
	for _, ele := range s {
		if st.Len() == 0 {
			st.PushBack(ele)
			continue
		}
		first := st.Back().Value.(rune)
		if isMatch(first, ele) {
			st.Remove(st.Back())
		} else {
			st.PushBack(ele)
		}
	}
	if st.Len() == 0 {
		return true
	}
	return false
}
func isMatch(first, sec rune) bool {
	if first == '(' && sec == ')' {
		return true
	}
	if first == '[' && sec == ']' {
		return true
	}
	if first == '{' && sec == '}' {
		return true
	}
	return false
}
