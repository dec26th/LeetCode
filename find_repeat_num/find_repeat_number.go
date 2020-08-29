package find_repeat_num

func findRepeatNumber(nums []int) int {
	numsMap := make(map[int]int, len(nums))

	for _, i := range nums {
		if _, ok := numsMap[i]; ok {
			return i
		}
		numsMap[i] = 1
	}
	return 0
}
