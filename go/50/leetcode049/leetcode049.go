package leetcode049

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	index := make(map[int]int, 0)
	var ret [][]string
	for _, str := range strs {
		key := 1
		for _, b := range str {
			key *= int(prime[b-'a'])
		}
		if val, ok := index[key]; ok {
			ret[val] = append(ret[val], str)
		} else {
			ret = append(ret, []string{str})
			index[key] = len(ret) - 1
		}
	}
	return ret
}

// 使用string做key
func groupAnagrams1(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	index := make(map[string]int, 0)
	var ret [][]string
	for _, str := range strs {
		bytes := make([]byte, 26)
		for _, b := range str {
			bytes[b-'a']++
		}
		key := string(bytes)
		if val, ok := index[key]; ok {
			ret[val] = append(ret[val], str)
		} else {
			ret = append(ret, []string{str})
			index[key] = len(ret) - 1
		}
	}
	return ret
}

func groupAnagrams2(strs []string) [][]string {
	var ret [][]string
	for _, str := range strs {
		isMatch := false
		for i := 0; i < len(ret); i++ {
			if match([]byte(str), []byte(ret[i][0])) {
				ret[i] = append(ret[i], str)
				isMatch = true
				break
			}
		}
		if !isMatch {
			ret = append(ret, []string{str})
		}
	}
	return ret
}
func match(str1, str2 []byte) bool {
	if len(str1) != len(str2) {
		return false
	}
	tmp := 0
	for i := 0; i < len(str1); i++ {
		for j := tmp; j < len(str2); j++ {
			if str1[i] == str2[j] {
				str2[tmp], str2[j] = str2[j], str2[tmp]
				tmp++
				break
			}
		}
	}
	return len(str1) == tmp
}
