package a1

// 一个整形数组中，每个数字都出现两次，只有一个数字出现一次，找出这个数字。
func solution(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// 一个整形数组中，每个数字都出现三次，只有一个数字出现一次，找出这个数字。
func solution1(nums []int) int {
	bits := make([]byte, 64)
	for i := 0; i < len(bits); i++ {
		bits[i] = 0
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(bits); j++ {
			// 记录这个位上是否为1，为1的话 bits数组就加1
			bits[j] += byte((nums[i] >> uint(j)) & 1)
		}
	}
	result := 0
	for j := 0; j < 64; j++ {
		if bits[j]%3 != 0 {
			// 把记录的二进制他转化成整形数字
			result += 1 << uint64(j)
		}
	}
	return result
}
