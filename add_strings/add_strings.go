package add_strings

import (
	"math"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	lenOfNum1 := len(num1)
	lenOfNum2 := len(num2)
	longer := int(math.Max(float64(lenOfNum1), float64(lenOfNum2)))

	carry := 0
	sum := 0
	for i := 0; i < longer; i++ {
		if i < lenOfNum1 {
			carry += int(num1[i] - '0')
		}
		if i < lenOfNum2 {
			carry += int(num2[i] - '0')
		}

		sum += (carry % 10) * int(math.Pow10(i))
		carry = carry / 10
	}
	sum += carry * int(math.Pow10(longer))

	return strconv.Itoa(sum)
}