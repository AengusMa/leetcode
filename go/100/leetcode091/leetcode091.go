package leetcode091

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	d1, d2 := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			d1 = 0
		}
		if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6' {
			d1 = d1 + d2
			d2 = d1 - d2
		} else {
			d2 = d1
		}
	}

	return d1
}
