package multiply_strings

import (
	"strconv"
)

func multiplyStrings(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	lenOfNum1 := len(num1)
	lenOfNum2 := len(num2)

	result := "0"
	carry := 0

	for i := lenOfNum1 - 1; i >= 0; i-- {

		char := num1[i]
		j := lenOfNum2 - 1
		var tempResult string
		for j >= 0 || carry != 0 {

			var temp int
			if j >= 0 {
				temp = int(char-'0')*int(num2[j]-'0') + carry
			} else {
				temp = carry
			}
			tempResult = strconv.Itoa(temp%10) + tempResult
			carry = temp / 10
			j--
		}
		for k := lenOfNum1 - 1; k > i; k-- {
			tempResult = tempResult + "0"
		}

		result = addStrings(result, tempResult)
	}
	return result
}

func addStrings(num1 string, num2 string) string {
	lenOfNum1 := len(num1)
	lenOfNum2 := len(num2)

	res := ""
	i, j, carry := lenOfNum1-1, lenOfNum2-1, 0
	var n1, n2 int

	for i >= 0 || j >= 0 || carry != 0 {

		if i >= 0 {
			n1 = int(num1[i] - '0')
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		} else {
			n2 = 0
		}

		temp := n1 + n2 + carry
		carry = temp / 10
		res = strconv.Itoa(temp%10) + res

		i--
		j--
	}
	return res
}

// todo 优化字符串相乘算法
