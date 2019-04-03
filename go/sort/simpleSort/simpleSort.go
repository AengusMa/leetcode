package simplesort

// SelectSort 选择排序
func SelectSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
			}
		}
	}
}

// InsertSort 插入排序
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > tmp; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		nums[j+1] = tmp
	}
}

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := true
		for j := 1; j < len(nums); j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
