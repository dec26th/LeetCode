package min_number

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	result := getMinNumberStringList(nums, 1)

	return strings.Join(result, "")
}

func getMinNumberStringList(nums []int, index int) []string {
	tempMap := make(map[int][]int, 10)
	for i := 0; i < len(nums); i++ {
		x := getFirstNum(nums[i], index)
		if _, ok := tempMap[x]; !ok {
			tempMap[x] = []int{
				nums[i],
			}
		} else {
			tempMap[x] = append(tempMap[x], nums[i])
		}
	}

	result := make([]string, len(nums))
	for i := 0; i < 10; i++ {
		if _, ok := tempMap[i]; !ok {
			continue
		} else {
			if len(tempMap[i]) == 1 {
				result = append(result, strconv.Itoa(tempMap[i][0]))
			} else {
				result = append(result, getMinNumberStringList(tempMap[i], index+1)...)
			}
		}
	}
	return result
}

func getFirstNum(num int, index int) int {
	nums := strconv.Itoa(num)
	if len(nums)-1 < index {
		return int(nums[len(nums)-1] - '0')
	} else {
		return int(nums[index-1] - '0')
	}
}
