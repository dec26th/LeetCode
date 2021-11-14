package missing_number

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left++
		} else {
			right--
		}
	}
	return left
}
