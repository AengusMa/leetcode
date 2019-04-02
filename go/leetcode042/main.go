package main

func trap(height []int) int {
	result := 0
	lh := 0
	for i := 0; i < len(height); i++ {
		if lh == 0 {
			if height[i] != 0 {
				lh = height[i]
			}
		} else {
			if height[i] < lh {
				result += lh - height[i]
			} else if height[i] > lh {

			}
		}

	}

	return result
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trap(height))
}
