package leetcode093

import "strconv"

func restoreIpAddresses(s string) []string {
	var validateArr []string
	for i1 := 1; i1 <= len(s)-3 && i1 < 4; i1++ {
		for i2 := i1 + 1; i2 <= len(s)-2 && i2-i1 < 4; i2++ {
			for i3 := i2 + 1; i3 <= len(s)-1 && i3-i2 < 4; i3++ {
				// can't begin with "0"
				if (len(s[0:i1]) > 1 && s[0:i1][0:1] == "0") || (len(s[i1:i2]) > 1 && s[i1:i2][0:1] == "0") ||
					(len(s[i2:i3]) > 1 && s[i2:i3][0:1] == "0") || (len(s[i3:]) > 1 && s[i3:][0:1] == "0") {
					continue
				}
				intTmp1, _ := strconv.Atoi(s[0:i1])
				intTmp2, _ := strconv.Atoi(s[i1:i2])
				intTmp3, _ := strconv.Atoi(s[i2:i3])
				intTmp4, _ := strconv.Atoi(s[i3:])

				if intTmp1 < 256 && intTmp2 < 256 && intTmp3 < 256 && intTmp4 < 256 {
					validateArr = append(validateArr, s[0:i1]+"."+s[i1:i2]+"."+s[i2:i3]+"."+s[i3:])
				}
			}
		}
	}
	return validateArr
}
