package leetcode043

import (
	"testing"
)

func Test043(t *testing.T) {
	if multiply("123", "456") != "56088" {
		t.Error("错误")
	}
	if multiply("0", "0") != "0" {
		t.Error("错误")
	}
	if multiply("9", "9") != "81" {
		t.Error("错误")
	}
	// fmt.Println(multiply("99", "99"))
	if multiply("99", "99") != "9801" {
		t.Error("错误")
	}
}
