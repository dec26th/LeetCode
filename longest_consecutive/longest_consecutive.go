package longest_consecutive

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	longestLength := 1

	numsMap := make(map[int]int, len(nums))
	for index, v := range nums {
		numsMap[v] = index
	}

	for v := range numsMap {
		temLength := 0
		findSmaller(v, &temLength, numsMap)
		findBigger(v+1, &temLength, numsMap)
		if temLength > longestLength {
			longestLength = temLength
		}
	}
	return longestLength
}

func findSmaller(target int, tempLength *int, numsMap map[int]int) {
	for {
		if _, ok := numsMap[target]; ok {
			delete(numsMap, target)
			*tempLength++
			target--

		} else {
			break
		}
	}
}

func findBigger(target int, tempLength *int, numsMap map[int]int) {
	for {
		if _, ok := numsMap[target]; ok {
			*tempLength++
			delete(numsMap, target)
			target++

		} else {
			break
		}
	}
}
