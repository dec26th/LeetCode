package exchange

func exchange(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if isEven(nums[i]) {
			result[right] = nums[i]
			right --
			continue
		}

		if isOdd(nums[i]) {
			result[left] = nums[i]
			left++
			continue
		}
	}

	return result
}

func isEven(num int) bool {
	return (num & 1) == 0
}

func isOdd(num int) bool {
	return (num & 1) == 1
}