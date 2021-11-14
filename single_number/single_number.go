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

// state machine
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

// 使用分组的方法，将两个不同的数字分别分到两个不同的组，然后分组在异或得到所需要的结果
// 首先将数组内的所有数字进行异或。得到一个结果，在该结果中，我们可以分析出，位为1的部位是两个只出现一次数字的不同的部位。那么可以通过这个不同的部位
// 1， 将整个数组分成两个不同的部分，然后在分别异或得到结果。
func twoNumberOnceDivide(nums []int) []int {
	var temp, sort1, sort2 int
	for i := 0; i < len(nums); i++ {
		temp ^= nums[i]
	}

	offset := 0
	for {
		if (temp & 1) == 0 {
			offset++
			temp = temp >> 1
		} else {
			break
		}
	}
	offset = 1 << offset
	for i := 0; i < len(nums); i++ {
		if (nums[i] & offset) == 0 {
			sort1 ^= nums[i]
		} else {
			sort2 ^= nums[i]
		}
	}
	return []int{sort1, sort2}
}
