package leetcode044

import "testing"

func Test044(t *testing.T) {
	if isMatch("aa", "a") != false {
		t.Error("错误")
	}
	if isMatch("aa", "*") != true {
		t.Error("错误")
	}
	if isMatch("cb", "?a") != false {
		t.Error("错误")
	}
	if isMatch("adceb", "*a*b") != true {
		t.Error("错误")
	}
	if isMatch("acdcb", "a*c?b") != false {
		t.Error("错误")
	}
}
