package leetcode043

func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))
	carry := 0
	for i := 0; i < len(num1)+len(num2); i++ {
		res[i] = '0'
	}
	for i := len(num2) - 1; i >= 0; i-- {
		resIndex := len(res) - (len(num2) - i)
		for j := len(num1) - 1; j >= 0; j-- {
			tmp := (num2[i]-'0')*(num1[j]-'0') + uint8(carry) + res[resIndex] - '0'
			res[resIndex] = tmp%10 + '0'
			if tmp >= 10 {
				carry = int(tmp) / 10
			} else {
				carry = 0
			}
			resIndex--
		}
		if carry != 0 {
			res[resIndex] = uint8(carry) + '0'
		}
		carry = 0
	}
	// 去掉开头的0
	valid := 0
	for i := 0; i < len(res); i++ {
		if res[i] == '0' && i != len(res)-1 {
			valid++
		} else {
			break
		}
	}
	return string(res[valid:])
}
