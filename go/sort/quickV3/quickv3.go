package quickv3

import "sync"

var wg sync.WaitGroup

// Sort 快速排序
func Sort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
	wg.Wait()
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	posL, posR := partition(nums, left, right)
	if right-left < 50 {
		wg.Add(1)
		go insertSort(nums, left, right)
	} else {
		quickSort(nums, left, posL)
		quickSort(nums, posR, right)
	}
}

// InsertSort 插入排序
func insertSort(nums []int, left, right int) {
	defer wg.Done()
	for i := left + 1; i <= right; i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > tmp; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		nums[j+1] = tmp
	}
}

func partition(arr []int, left, right int) (int, int) {
	pIndex := choosePivotMedianOfThree(arr, left, right)
	if pIndex != left {
		arr[pIndex], arr[left] = arr[left], arr[pIndex]
	}
	p := arr[left]
	i := left + 1
	sameLen := 1
	for j := left + 1; j <= right; j++ {
		if arr[j] <= p {
			arr[j], arr[i] = arr[i], arr[j]
			// 相等的元素放在头部
			if arr[i] == p {
				arr[i], arr[left+sameLen] = arr[left+sameLen], arr[i]
				sameLen++
			}
			i++

		}
	}
	posRight := i
	posLeft := posRight - sameLen - 1
	// 相等的元素移动中间
	for ; sameLen > 0 && arr[i-1] != p; sameLen-- {
		arr[i-1], arr[left+sameLen-1] = arr[left+sameLen-1], arr[i-1]
		i--
	}
	return posLeft, posRight
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
