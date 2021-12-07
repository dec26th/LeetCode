package search

func search(nums []int, target int) int {
	index := -1

	left := 0
	right := len(nums) - 1

	for left < right-1 {
		mid := (left + right) / 2
		leftValue, midValue, rightValue := nums[left], nums[mid], nums[right]

		switch {

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

func searchNum(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 1
		}
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] < target {
			left++
		}
		if nums[right] > target {
			right--
		}
		if nums[left] == target && nums[right] == target {
			return right - left + 1
		}
	}
	return 0
}

func searchToday(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (right + left) / 2

		if nums[middle] == target {
			return middle
		}
		if left == right {
			return -1
		}

		if nums[left] < nums[right] {
			if nums[middle] < target {
				left = middle + 1
				continue
			} else if nums[middle] > target {
				right = middle - 1
				continue
			}
		}

		if nums[left] > nums[right] {
			if target < nums[right] {
				left = middle
			}

			if target > nums[left] {
				right = middle
			}

			if target >= nums[right] && target <= nums[len(nums)-1] {
				left = right
				right = len(nums) - 1
			}

			if target <= nums[left] && target >= nums[0] {
				right = left
				left = 0
			}

		}
		if left+1 == right {
			return -1
		}
	}

	return -1
}
