package min_path_sum

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid);i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i - 1 >= 0 && j - 1 >= 0 {
				dp[i][j] = min(dp[i - 1][j] + grid[i][j], dp[i][j - 1] + grid[i][j])
				continue
			}

			if i - 1 >= 0 {
				dp[i][j] = dp[i - 1][j] + grid[i][j]
				continue
			}

			if j - 1 >= 0 {
				dp[i][j] = dp[i][j - 1] + grid[i][j]
				continue
			}
		}
	}

	return dp[len(grid) - 1][len(grid[0]) - 1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
