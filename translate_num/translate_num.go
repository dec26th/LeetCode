package translate_num

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	if num < 26 {
		return 2
	}
	if num < 100 {
		return 1
	}
	nums := getNumList(num)
	result := make([]int, len(nums))
	result[0] = 1
	if (nums[len(nums)-1]*10+nums[len(nums)-2]) < 26 && nums[len(nums)-1] != 0 {
		result[1] = 2
	} else {
		result[1] = 1
	}
	for j := len(nums) - 3; j >= 0; j-- {
		if (nums[j+1]*10+nums[j]) < 26 && nums[j+1] != 0 {
			result[len(nums)-1-j] = result[len(nums)-2-j] + result[len(nums)-3-j]
		} else {
			result[len(nums)-1-j] = result[len(nums)-2-j]
		}
	}
	return result[len(result)-1]
}

func getNumList(num int) []int {
	nums := []int{}
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	return nums
}
