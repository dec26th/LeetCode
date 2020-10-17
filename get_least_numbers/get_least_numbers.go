package get_least_numbers


func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	quickSearch(arr, 0, len(arr), k)
	return arr[:k]
}

func quickSearch(arr []int, start, end, k int) {

	index := partition(arr, start, end)
	if index == k - 1 {
		return
	}
	if index > k - 1 {
		quickSearch(arr, start, index, k)
	} else {
		quickSearch(arr, index + 1, end, k)
	}
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i < right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index += 1
		}
	}

	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}