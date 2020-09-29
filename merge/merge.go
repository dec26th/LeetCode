package merge

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	result := *new([][]int)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lenOfResult := len(result)
		for index, value := range result {
			switch  {
			case intervals[i][0] > value[1]:
				if lenOfResult == index + 1 {
					result = append(result, intervals[i])
				}
			case intervals[i][1] < value[1]:
			case intervals[i][1] >= value[1]:
				result[index][1] = intervals[i][1]
			}
		}
	}

	return result
}
