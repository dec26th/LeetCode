package search

func search(nums []int, target int) int {
	index := -1

	left := 0
	right := len(nums) - 1

	for left < right - 1 {
		mid := (left + right) / 2
		leftValue, midValue, rightValue := nums[left], nums[mid], nums[right]

		switch  {

		case midValue == target:
			 return mid

		case leftValue <= target && target <= midValue:
			 right = mid

		case target >= midValue && target <= rightValue:
			left = mid + 1

		case leftValue > midValue:
			right = mid

		case midValue > rightValue:
			left = mid

		default:
			return -1
		}

	}

	if nums[left] == target {
		index = left
	} else if nums[right] == target {
		index = right
	}

	return index
}
