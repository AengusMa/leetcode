package leetcode075

func sortColors(nums []int) {
	red := 0
	white := 0
	blue := len(nums) - 1
	for white <= blue {
		if nums[white] == 0 {
			nums[red], nums[white] = nums[white], nums[red]
			white++
			red++
		} else if nums[white] == 1 {
			white++
		} else {
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}
