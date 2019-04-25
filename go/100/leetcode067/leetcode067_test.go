package leetcode067

import "testing"

func Test067(t *testing.T) {
	if addBinary("11", "1") != "100" {
		t.Error("error")
	}
	if addBinary("1010", "1011") != "10101 v" {
		t.Error("error")
	}
}
