package leetcode017

var data map[string]string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	data = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	var result []string
	for index, num := range digits {
		tmp := data[string(num)]
		l := len(result)
		for i := 0; i < len(tmp); i++ {
			if index == 0 {
				result = append(result, string(tmp[i]))
			} else {
				if i == 0 {
					for k := 0; k < len(result); k++ {
						result[k] += string(tmp[i])
					}
				} else {
					for k := 0; k < l; k++ {
						t := result[k][:len(result[k])-1] + string(tmp[i])
						result = append(result, t)
					}
				}
			}
		}
	}
	return result
}
