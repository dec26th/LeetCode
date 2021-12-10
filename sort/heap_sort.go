package sort

func HeapSort(nums []int) []int {
	lenOfNums := len(nums)
	buildMaxHeap(nums, lenOfNums)

	for i := lenOfNums - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lenOfNums--
		heapify(nums, 0, lenOfNums)
	}
	return nums
}

func buildMaxHeap(nums []int, len int) {
	for i := len / 2; i >= 0; i-- {
		heapify(nums, i, len)
	}
}

func heapify(nums []int, i int, len int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < len && nums[left] > nums[largest] {
		largest = left
	}

	if right < len && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, len)
	}
}

func MyHeapSort(nums []int) []int {
	_buildMaxHeap(nums)

	lengthOfNums := len(nums)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		lengthOfNums--
		_heapify(nums, 0, lengthOfNums)
	}

	return nums
}

func _buildMaxHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		_heapify(nums, i, len(nums))
	}
}

func _heapify(nums []int, i int, length int) {
	left := 2*i + 1
	right := 2*i + 2

	origin := nums[i]
	if left < length && nums[left] > nums[i] {
		nums[i], nums[left] = nums[left], nums[i]
	}

	if right < length && nums[right] > nums[i] {
		nums[i], nums[right] = nums[right], nums[i]
	}
	if origin != nums[i] {
		_heapify(nums, left, length)
		_heapify(nums, right, length)
	}

}
