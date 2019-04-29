package leetcode079

import (
	"testing"
)

func Test079(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	if exist(board, "ABCCED") != true {
		t.Error("error")
	}
	if exist(board, "SEE") != true {
		t.Error("error")
	}
	if exist(board, "ABCB") != false {
		t.Error("error")
	}
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	if exist(board, "ABCESEEEFS") != true {
		t.Error("error")
	}
	board = [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}
	if exist(board, "aaaaaaaaaaaaa") != false {
		t.Error("error")
	}
}

func TestYu(t *testing.T) {
	num := 111
	num ^= 255
	println(num)
	num ^= 100
	println(num)
}
