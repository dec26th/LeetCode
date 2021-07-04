package print_numbers

import "math"

func printNumbers(n int) []int {
	length := int(math.Pow10(n)) - 1

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = i + 1
	}
	return result
}