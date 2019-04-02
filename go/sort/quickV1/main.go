package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pos := partition(nums, left, right)
	quickSort(nums, left, pos-1)
	quickSort(nums, pos+1, right)
}
func partition(nums []int, left, right int) int {
	p_index := left
	if p_index != left {
		nums[p_index], nums[left] = nums[left], nums[p_index]
	}
	i := left + 1
	p := nums[left]
	for j := left + 1; j <= right; j++ {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[left], nums[i-1] = nums[i-1], nums[left]
	return i - 1
}
func choosePivotMedianOfThree(arr []int, left, right int) int {
	mid := left + ((right - left) >> 1)
	if (arr[mid] > arr[right] && arr[mid] < arr[left]) || (arr[mid] > arr[left] && arr[mid] < arr[right]) {
		return mid
	} else if (arr[left] > arr[right] && arr[left] < arr[mid]) || (arr[left] > arr[mid] && arr[left] < arr[right]) {
		return left
	} else {
		return right
	}
}
func main() {
	nums := getArray(90, 10000)
	fmt.Println("原始：", nums)
	qSort(nums)
	fmt.Println(nums)

}

func getArray(len, max int) []int {
	rand.Seed(time.Now().Unix())
	result := []int{}
	for i := 0; i < len; i++ {
		result = append(result, rand.Intn(max))
	}
	return result
}
