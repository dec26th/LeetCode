package sort

func HeapSort(nums []int) []int {
	lenOfNums := len(nums)
	buildMaxHeap(nums, lenOfNums)

	for i := lenOfNums - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lenOfNums --
		heapify(nums, 0, lenOfNums)
	}
	return nums
}


func buildMaxHeap(nums []int, len int) []int {
	for i := len / 2; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, len)
 	}
	return nums
}

func heapify(nums []int, i int, len int) {
	left := 2 * i + 1
	right := 2 * i + 2
	largest := 1
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