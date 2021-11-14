package find_length_of_LCIS

func findLengthOfLCIS(nums []int) int {
	maxLength := 1
	lenOfNums := len(nums)
	if lenOfNums == 0 {
		return 0
	}
	dp := make([]int, lenOfNums)

	for i := 0; i < lenOfNums-1; i++ {
		if dp[i] == 0 {
			dp[i] = 1
		}

		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
			maxLength = max(dp[i+1], maxLength)
		}
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
