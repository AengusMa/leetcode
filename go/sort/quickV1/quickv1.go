package quickv1

// Sort 快速排序
func Sort(nums []int) {
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
	pIndex := left
	if pIndex != left {
		nums[pIndex], nums[left] = nums[left], nums[pIndex]
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
