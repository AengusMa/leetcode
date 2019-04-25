package leetcode070

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		next := first + second
		first = second
		second = next
	}
	return second
}

func climbStairs1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}
