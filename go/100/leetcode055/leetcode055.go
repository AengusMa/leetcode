package leetcode055

//  [3,2,1,0,4]
//   0 1 2 3 4
// 1, 1, 0, 1
// 0  1  2  3
// 0  1  2  4
func canJump(nums []int) bool {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if i > pos {
			return false
		}
		if nums[i]+i > pos {
			pos = nums[i] + i
		}
	}
	return true
}
func canJump1(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	maxJump := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]+i > maxJump {
			maxJump = nums[i] + i
		}
		if nums[i] == 0 && i >= maxJump {
			return false
		}
	}
	return true
}

func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	ret := true
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] == 0 {
			canSkip := false
			for i := 0; i < j; i++ {
				if j < nums[i]+i {
					canSkip = true
				}
			}
			if !canSkip {
				ret = false
			}
		}
	}
	return ret
}
