package leetcode012

func intToRoman(num int) string {
	result := ""
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	values := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	for i := len(values) - 1; i >= 0; i-- {
		m := num / values[i]
		num = num % values[i]
		for ; m != 0; m-- {
			result += symbols[i]
		}
	}
	return result
}

func intToRoman1(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[num%1000/100] + X[num%100/10] + I[num%10]
}
