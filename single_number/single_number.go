package single_number

func singleNumber(nums []int) int {
	record := make(map[int]int)

	for _, num := range nums {
		if _, ok := record[num]; ok {
			record[num] += 1
		} else {
			record[num] = 1
		}
	}

	for k, v := range record {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberBit(nums []int) int {
	var result int
	for _, value := range nums {
		result ^= value
	}
	return result
}

func singleNumberOtherThree(nums []int) int {
	record := make(map[int]int)

	for _, num := range nums {
		if _, ok := record[num]; ok {
			record[num] += 1
		} else {
			record[num] = 1
		}
	}

	for k, v := range record {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberOhterThreeBit(nums []int) int {
	one, two := 0, 0

	for _, num := range nums {
		one = one ^ num&(^two)
		two = two ^ num&(^one)
	}
	return one
}

func twoNumberOnce(nums []int) []int {
	record := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums {
		if _, ok := record[num]; ok {
			record[num] += 1
		} else {
			record[num] = 1
		}
	}

	for k, v := range record {
		if v == 1 {
			result = append(result, k)
			if len(result) == 2 {
				return result
			}
		}
	}
	return []int{0, 0}
}

func twoNumberOnceBit(nums []int) []int {
	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}

	diff := bitmask & (-bitmask)
	var x int
	for _, v := range nums {
		if (v & diff) != 0 {
			x ^= v
		}
	}

	return []int{x, bitmask ^ x}
}
