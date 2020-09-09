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