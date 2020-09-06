package three_sum

import (
	"sort"
)
// 超时
func threeSum(nums []int) [][]int {
	lenOfNums := len(nums)
	result := make([][]int, 0)
	sort.Ints(nums)

	for first := 0; first < lenOfNums; first ++ {

		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := lenOfNums - 1
		target := - nums[first]

		for second := first + 1; second < third; second ++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			if nums[second] + nums[third] == target {
				result = append(result, []int {nums[first], nums[second], nums[third]})
			}
			if nums[second] + nums[third] > target {
				for j := third; j > second ; j-- {

					if j < third - 1 && nums[j] == nums[j + 1] {
						continue
					}

					if nums[second] + nums[j] == target {
						result = append(result, []int {nums[first], nums[second], nums[j]})
					}

					if nums[second] + nums[j] < target {
						break
					}
				}
			}
			if nums[second] + nums[third] < target {
				continue
			}
		}

	}
	return result
}

func ThreeSumImproved(nums []int) [][]int {
	lenOfNums := len(nums)
	result := make([][]int, 0)
	sort.Ints(nums)

	for first := 0; first < lenOfNums; first ++ {

		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := lenOfNums - 1
		target := - nums[first]

		for second := first + 1; second < third; second ++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}

			for second < third && nums[second] + nums[third] > target {
				third --
			}

			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
				result = append(result, []int {nums[first], nums[second], nums[third]})
			}

		}

	}
	return result
}
