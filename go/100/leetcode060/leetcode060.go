package leetcode060

func getPermutation1(n int, k int) string {
	// f保存i!的值
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = 1 * f[i-1]
	}
	ret := make([]byte, n)
	nums := "123456789"
	for i := n; i >= 1; i-- {
		j := k / f[i-1]
		k %= f[i-1]
		ret[i-1] = nums[j]
	}
	return string(ret)
}

func getPermutation(n int, k int) string {
	ret := make([]byte, 0)
	for i := 1; i <= n; i++ {
		ret = append(ret, byte('1'+i-1))
	}
	if k == 1 {
		return string(ret)
	}
	for i := 2; i <= k; i++ {
		nextPermutation(ret)
	}
	return string(ret)
}

func nextPermutation(nums []byte) {
	tmp := -1
	// 倒序找到第一个非递增的数
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			tmp = i
			break
		}
	}
	// 如果没有翻转数组
	if tmp == -1 {
		reverseSort(nums)
		return
	}
	for j := len(nums) - 1; j > tmp; j-- {
		if nums[j] > nums[tmp] {
			nums[j], nums[tmp] = nums[tmp], nums[j]
			reverseSort(nums[tmp+1:])
			break
		}
	}
}
func reverseSort(nums []byte) {
	start := 0
	end := len(nums) - 1
	for end > start {
		nums[end], nums[start] = nums[start], nums[end]
		end--
		start++
	}
}
