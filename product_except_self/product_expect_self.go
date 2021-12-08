package product_except_self


func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	first, second := 1, 1

	for i := 0; i < len(nums); i++ {
		result[i] = first
		first *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= second
		second *= nums[i]
	}

	return result
}