package shellsort

// Sort 希尔排序
func Sort(nums []int) {
	for step := len(nums) / 2; step > 0; step = step / 2 {
		// 从增量那组开始进行插入排序，直至完毕
		for i := step; i < len(nums); i++ {
			for j := i; j-step >= 0 && nums[j-step] > nums[j]; j -= step {
				nums[j], nums[j-step] = nums[j-step], nums[j]
			}
		}
	}
}
