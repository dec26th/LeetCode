package longest_consecutive

var numsMap map[int]bool

func longestConsecutive(nums []int) int {
	longestLength := 0

	numsMap = make(map[int]bool, len(nums))
	for _, v := range nums {
		numsMap[v] = true
	}

	for _, v := range nums {
		temLength := 0
		if res, ok := numsMap[v]; ok && res {
			findSmaller(v, &temLength)
			findBigger(v + 1, &temLength)
			if temLength > longestLength {
				longestLength = temLength
			}
		}
	}
	return longestLength
}

func findSmaller(target int, tempLength *int) {
	for {
		if res, ok := numsMap[target]; ok && res {
			*tempLength ++
			numsMap[target] = false
			target--

		} else {
			break
		}
	}
}

func findBigger(target int, tempLength *int) {
	for {
		if res, ok := numsMap[target]; ok && res {
			*tempLength ++
			numsMap[target] = false
			target++

		} else {
			break
		}
	}
}
