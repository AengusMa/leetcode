package utils

import "math/rand"

// GetSortArray 获取有序数组
func GetSortArray(len int, desc bool) []int {
	result := []int{}
	if desc {
		for i := len - 1; i > 0; i-- {
			result = append(result, i)
		}
	} else {
		for i := 0; i < len; i++ {
			result = append(result, i)
		}
	}
	return result
}

// GetArray 获取无序数组
func GetArray(len, max int) []int {
	result := []int{}
	for i := 0; i < len; i++ {
		result = append(result, rand.Intn(max))
	}
	return result
}

// GetSameArray 获取元素重复数组
func GetSameArray(len int) []int {
	result := []int{}
	for i := 0; i < len; i++ {
		result = append(result, 10)
	}
	return result
}
