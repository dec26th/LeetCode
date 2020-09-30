package minimum_total

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 && len(triangle[0]) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle))
	minimum := math.MaxInt32

	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j]
		}
	}

	for i := 1; i < len(dp) - 1; i++ {
		dp[i][0] = dp[i - 1][0] + dp[i][0]
		length := len(dp[i])
		dp[i][length - 1] = dp[i - 1][length - 2] + dp[i][length - 1]

		for j := 1; j < length - 1; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i - 1][j - 1]) + dp[i][j]
		}
	}

	lenOfDP := len(dp) - 1
	length := len(dp[lenOfDP])
	dp[lenOfDP][0] = dp[lenOfDP - 1][0] + dp[lenOfDP][0]
	if dp[lenOfDP][0] < minimum {
		minimum = dp[lenOfDP][0]
	}
	dp[lenOfDP][length - 1] = dp[lenOfDP - 1][length - 2] + dp[lenOfDP][length - 1]
	if dp[lenOfDP][length - 1] < minimum {
		minimum = dp[lenOfDP][length - 1]
	}

	for j := 1; j  < length - 1; j++ {
		dp[lenOfDP][j] = min(dp[lenOfDP - 1][j - 1], dp[lenOfDP - 1][j]) + dp[lenOfDP][j]
		if dp[lenOfDP][j] < minimum {
			minimum = dp[lenOfDP][j]
		}
	}
	return minimum
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}