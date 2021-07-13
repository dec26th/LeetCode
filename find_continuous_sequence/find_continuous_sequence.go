package find_continuous_sequence

func findContinuousSequence(target int) [][]int {
	start, end := 1, 2
	result := make([][]int, 0)

	for start != end {
		if sum(start, end) == target {
			result = append(result, getSlice(start, end))
			start++
		}
		if sum(start, end) < target {
			end++
		}
		if sum(start, end) > target {
			start++
		}
	}
	return result
}

func sum(start, end int)int {
	return (start + end) * (end - start + 1) / 2
}

func getSlice(start, end int) []int {
	result := make([]int, end - start + 1)
	for i := 0; i < len(result); i++ {
		result[i] = start + i
	}
	return result
}
