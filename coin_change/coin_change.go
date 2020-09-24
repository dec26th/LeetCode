package coin_change

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		for index, coin := range coins {
			if index == 0 {
				dp[i] = amount + 1
			}
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
