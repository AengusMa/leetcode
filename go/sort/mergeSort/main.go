package main

import "fmt"

func sort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		// 将两个有序数组合并
		merge(nums, left, mid, right)

	}
}
func merge(nums []int, left, mid, right int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	i := mid + 1
	j := left
	for left <= mid && i <= right {
		if tmp[left] < tmp[i] {
			nums[j] = tmp[left]
			left++
		} else {
			nums[j] = tmp[i]
			i++
		}
		j++
	}
	for ; left <= mid; left++ {
		nums[j] = tmp[left]
		j++
	}
	for ; i <= right; i++ {
		nums[j] = tmp[i]
		j++
	}
}

func main() {
	// nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 14, 4, 4, 1, 1, 1, 4, 4, 4, 4, 4, 2, 5, 4, 4, 7, 4, 4, 4, 4, 4, 3, 1, 1, 1, 1, 1}
	nums := []int{8, 15, 4, 15, 11, 2, 13, 12, 4, 14, 0, 10, 6, 18, 9, 15, 6, 13, 12, 14}
	// nums := []int{4, 4, 2, 0}
	fmt.Println(nums)
	sort(nums)
	fmt.Println(nums)
}
