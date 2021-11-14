package min_array

import "sort"

// 剑指 Offer 11. 旋转数组的最小数字

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，
// 输出旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

func minArray(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func minArrayBiSearch(numbers []int) int {
	low := 0
	high := len(numbers) - 1

	for low < high {
		mid := (high-low)/2 + low
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
