package search_range

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums) - 1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			left = middle - 1
			for left >= 0 && nums[left] == target {
				left--
			}
			right = middle + 1
			for right < len(nums) && nums[right] == target {
				right++
			}

			return []int{left+1, right-1}
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return []int{-1, -1}
}