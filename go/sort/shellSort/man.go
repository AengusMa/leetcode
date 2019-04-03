package main

func sort(nums []int) {
	if len(nums) == 1 {
		return
	}
	for step := len(nums) / 2; step > 0; step = step / 2 {
		//从增量那组开始进行插入排序，直至完毕
		for i := step; i < len(nums); i++ {
			// j := i
			// temp := nums[j]
			// for j-step >= 0 && nums[j-step] > temp {
			// 	nums[j] = nums[j-step]
			// 	j = j - step
			// }
			// nums[j] = temp
			print("step:", step, " ")
			for j := i; j-step >= 0; j -= step {

				print(nums[j], " ")
			}
			println()
			for j := i; j-step >= 0 && nums[j-step] > nums[j]; j -= step {
				nums[j], nums[j-step] = nums[j-step], nums[j]
			}
			// fmt.Println(nums)
		}
	}
}
func main() {
	// nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 14, 4, 4, 1, 1, 1, 4, 4, 4, 4, 4, 2, 5, 4, 4, 7, 4, 4, 4, 4, 4, 3, 1, 1, 1, 1, 1}
	nums := []int{8, 15, 4, 15, 11, 2, 13, 12, 4, 14, 0, 10, 6, 18, 9, 15, 6, 13, 12, 14}
	// nums := []int{4, 4, 2, 0}
	// fmt.Println(nums)
	sort(nums)
	// fmt.Println(nums)
}
