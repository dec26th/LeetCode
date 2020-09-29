package sort

func QuickSort(nums []int) []int {
	return _quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, left int, right int) []int {
	if left < right {
		partitionIndex := partition(nums, left, right)
		_quickSort(nums, left, partitionIndex-1)
		_quickSort(nums, partitionIndex+1, right)
	}
	return nums
}

func partition(nums []int, left int, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index += 1
		}
	}

	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
