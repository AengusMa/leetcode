package leetcode060

import (
	"testing"
)

func Test060(t *testing.T) {
	// "123"
	// "132"
	// "213"
	// "231"
	// "312"
	// "321"
	if getPermutation(3, 3) != "213" {
		t.Error("error")
	}
	if getPermutation(4, 9) != "2314" {
		t.Error("error")
	}
}
