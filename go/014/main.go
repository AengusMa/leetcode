package main

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result = commonPrefix(strs[i], result)
	}
	return result
}
func commonPrefix(str1, str2 string) string {
	result := ""
	minLen := min(len(str1), len(str2))
	for i := 0; i < minLen; i++ {
		if str1[i] == str2[i] {
			result += string(str1[i])
		} else {
			break
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	strs := []string{"caa", ""}
	println(longestCommonPrefix(strs))
}
