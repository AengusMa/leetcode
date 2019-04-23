package leetcode029

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}
	a, b := abs(dividend), abs(divisor)
	res := 0
	for a >= b {
		a -= b
		res++
	}
	if divisor > 0 && dividend < 0 {
		return -res
	}
	if divisor < 0 && dividend > 0 {
		return -res
	}
	if res == 2147483648 {
		return res - 1
	}
	return res
}

func abs(num int) int64 {
	tmp := int64(num)
	if num > 0 {
		return tmp
	}
	return -tmp
}
