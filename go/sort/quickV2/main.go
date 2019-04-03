package quickv2

// Sort 优化重复数据多的情况
func Sort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	// 数据分为三个区
	posL, posR := partition(arr, low, high)
	quickSort(arr, low, posL)
	quickSort(arr, posR, high)

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
func choosePivotFirst(arr []int, left, right int) int {
	return left
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
