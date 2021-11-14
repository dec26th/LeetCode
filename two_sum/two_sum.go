package two_sum

func twoSum(nums []int, target int) []int {
	head, tail := 0, len(nums)-1

	for head != tail {
		if nums[head]+nums[tail] < target {
			head++
		} else if nums[head]+nums[tail] > target {
			tail--
		} else {
			return []int{nums[head], nums[tail]}
		}
	}

	return []int{0, 0}
}
