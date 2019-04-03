package mergesort

// Sort 归并排序
func Sort(nums []int) {
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
