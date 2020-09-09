package add_strings

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	lenOfNum1 := len(num1)
	lenOfNum2 := len(num2)

	res := ""
	i, j, carry := lenOfNum1 - 1, lenOfNum2 - 1, 0
	var n1, n2 int

	for i >= 0 || j >= 0 || carry != 0 {

		if i >= 0 {
			n1 = int(num1[i] - '0')
		}else {
			n1 = 0
		}
		if  j >= 0 {
			n2 = int(num2[j] - '0')
		}else {
			n2 = 0
		}

		temp := n1 + n2 + carry
		carry = temp / 10
		res = strconv.Itoa(temp % 10) + res

		i --
		j --
	}
	return res
}