package find_nth_digit

import "strconv"

func findNthDigit(n int) int {
	var result int
	x, y := 1, 9

	for n > x*y {
		n -= x * y
		x++
		y *= 10
	}

	ten := getTen(x - 1)

	z := (n - 1) % x
	n = (n - 1) / x

	temp := ten + n

	tempStr := strconv.Itoa(temp)
	num := tempStr[z]

	result, _ = strconv.Atoi(string(num))
	return result
}

func getTen(n int) int {
	result := 1
	for n != 0 {
		result *= 10
		n--
	}

	return result
}
