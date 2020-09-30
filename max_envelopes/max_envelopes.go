package max_envelopes

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 || len(envelopes) == 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
	}

	maxNum := 1
	for i := 0; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[j] + 1, dp[i])
				maxNum = max(dp[i], maxNum)
			}
		}
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
