package main

import "fmt"

func sort(nums []int) {
	selectSort(nums, 0, len(nums)-1)
}
func sort1(nums []int) {
	bubbleSort(nums, 0, len(nums)-1)
}
func sort2(nums []int) {
	insertSort(nums, 0, len(nums)-1)
}
func selectSort(nums []int, left, right int) {
	for i := left + 1; i < right; i++ {
		for j := i; j <= right; j++ {
			if nums[j] < nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
			}
		}
		fmt.Println(nums)
	}
}
func insertSort(nums []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= left && nums[j] > tmp; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		nums[j+1] = tmp
		fmt.Println(nums)
	}
}
func bubbleSort(nums []int, left, right int) {
	for i := left; i <= right; i++ {
		flag := true
		for j := left + 1; j <= right; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
		fmt.Println(nums)
	}
}

func main() {
	// nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 14, 4, 4, 1, 1, 1, 4, 4, 4, 4, 4, 2, 5, 4, 4, 7, 4, 4, 4, 4, 4, 3, 1, 1, 1, 1, 1}
	nums := []int{8, 15, 4, 15, 11, 2, 13, 12, 4, 14, 0, 10, 6, 18, 9, 15, 6, 13, 12, 14}
	// nums := []int{4, 4, 2, 0}
	fmt.Println(nums)
	sort2(nums)
	fmt.Println(nums)
}
