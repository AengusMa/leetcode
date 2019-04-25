package leetcode067

// TODO
func addBinary(a string, b string) string {
	ret := make([]byte, 0)
	index1, index2 := len(a)-1, len(b)-1
	carry := byte('0')
	for ; index1 >= 0; index1-- {
		for ; index2 >= 0; index2-- {
			if a[index1] == '1' && b[index2] == '1' {
				ret = append(ret, carry)
				carry = byte('1')
			}
		}
	}
	return string(ret)
}
