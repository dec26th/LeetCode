package max_sub_array

import "math"

func maxSubArray(nums []int) int {
	var maxSum int
	lenOfNums := len(nums)
	for i := 0; i < lenOfNums; i++ {
		lastSum := 0
		for j := i; j < lenOfNums; j++ {
			if i == 0 && j == 0 {
				lastSum += nums[j]
				maxSum = nums[0]
			} else {
				lastSum += nums[j]
				maxSum = int(math.Max(float64(lastSum), float64(maxSum)))
			}
		}
	}

	return maxSum
}

func maxSubArrayDP(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ { // f(i) = max(f(i - 1) + ai, ai)
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
