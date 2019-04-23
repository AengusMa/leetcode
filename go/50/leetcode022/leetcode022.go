package leetcode022

func generateParenthesis(n int) []string {
	return add("", n, 0)
}
func add(str string, canOpen, canClose int) (result []string) {
	if canOpen == 0 && canClose == 0 {
		return []string{str}
	}
	if canOpen > 0 {
		result = append(result, add(str+"(", canOpen-1, canClose+1)...)
	}
	if canClose > 0 {
		result = append(result, add(str+")", canOpen, canClose-1)...)
	}
	return
}
