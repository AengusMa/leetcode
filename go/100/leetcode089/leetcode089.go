package leetcode089

import "fmt"

func grayCode(n int) []int {
	ret := make([]int, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		ret[i] = i ^ (i >> 1)
	}
	return ret
}

// (0,1) ---> (00 01,11 10) -->(000 001 011 010,110 111 101 100)
func grayCode1(n int) []int {
	ret := make([]int, 1)
	ret[0] = 0
	for i := 0; i < n; i++ {
		fmt.Println(ret, 1<<uint(i))
		for k := len(ret) - 1; k >= 0; k-- {
			ret = append(ret, ret[k]|1<<uint(i))
		}
	}
	return ret
}
