package majority_element

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	record := make(map[int]int, len(nums)/2+1)

	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = 1
			continue
		}

		if record[nums[i]] >= len(nums)/2 {
			return nums[i]
		}

		record[nums[i]] += 1
	}
	return 0
}
